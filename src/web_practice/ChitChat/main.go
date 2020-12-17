package main

import (
	"log"
	"net/http"
	"time"
)

const (
	address      string = "0.0.0.0:8080"
	WriteTimeout int64  = 600
	ReadTimeout  int64  = 10
)

func main() {

	log.Printf("[WEB-man](log) started at %s\n", address)

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:           address,
		Handler:        mux,
		ReadTimeout:    time.Duration(ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
