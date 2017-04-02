package main

import (
	"net/http"
	"log"
	"os"
	"os/signal"
	"syscall"
	"context"
)

func main() {
	config := mustParseConfig()
	mux := http.NewServeMux()
	s := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	db := newSqlStorage(&config)
	h := newHandler(db)

	mux.HandleFunc("/add-object", h.addObject)
	mux.HandleFunc("/get-object", h.getObject)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func(){
		<-signals
		s.Shutdown(context.TODO())
		db.db.Close()
		log.Print("Server stopped")
	}()

	log.Print("Server started")
	s.ListenAndServe()
}

