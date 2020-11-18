package main

import (
	"net/http"
)

func main() {
	//just a placeholder
}

func serv() *http.Server {
	server := http.Server{
		Addr:    ":8080",
		Handler: dummy{},
	}

	go server.ListenAndServe()
	return &server
}

type dummy struct{}

func (d dummy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("yes"))
}
