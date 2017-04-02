package main

import (
	"net/http"
	"encoding/json"
	"log"
)

type handler struct {
	storage storage
}

func newHandler(s storage) *handler {
	return &handler{storage: s}
}

func (h *handler) addObject(w http.ResponseWriter, r *http.Request) {
	var obj object

	err := json.NewDecoder(r.Body).Decode(&obj)
	if err != nil {
		log.Printf("Failed to decode payload: %s", err)
		http.Error(w, "Invlid payload", http.StatusBadRequest)
		return
	}

	err = h.storage.addObject(obj)
	if err != nil {
		log.Printf("Unexpected error: %s", err)
		http.Error(w, "Unexpected error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("OK"))
}

func (h *handler) getObject(w http.ResponseWriter, r *http.Request) {
	var in struct{ Id int }

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		log.Printf("Failed to decode payload: %s", err)
		http.Error(w, "Invlid payload", http.StatusBadRequest)
		return
	}

	obj, err := h.storage.getObject(in.Id)
	if err != nil {
		log.Printf("Unexpected error: %s", err)
		http.Error(w, "Unexpected error", http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(obj)
	if err != nil {
		log.Printf("Failed to marshall %+v: %s", obj, err)
		http.Error(w, "Unexpected error", http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}
