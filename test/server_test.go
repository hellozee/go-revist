package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := serv()
	resp, err := http.Get("http://localhost:8080")

	if err != nil {
		t.Fatal(err.Error())
	}

	if resp.StatusCode != 200 {
		t.Fatal("Status Code not 200")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err.Error())
	}

	if string(body) != "yes" {
		t.Fatal("Body malformed")
	}

	server.Shutdown(context.Background())
}

func TestBadInput(t *testing.T) {
	server := serv()
	_, err := http.Get("http://localhost")

	if err == nil {
		t.Error("No server should ne running on port 80")
	}

	server.Shutdown(context.Background())
}
