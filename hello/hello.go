package main

import (
	"net/http"
	"io"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func hello2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "What's up?")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", hello2)

	go http.ListenAndServe(":8000", mux)
	go http.ListenAndServe(":8001", mux2)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("awaiting signal")
	<-sigs
	fmt.Println("exiting")
}
