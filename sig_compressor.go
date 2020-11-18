package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var nameHolder []string

func genNames() {
	for {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		nameHolder = append(nameHolder, "yes")
	}
}

func main() {
	go serv()
	go genNames()

	for {
		time.Sleep(7 * time.Second)
		newSlice := append([]string(nil), nameHolder...)
		nameHolder = nameHolder[:0]
		doRequest(newSlice...)
	}
}

type Names []string

func serv() {
	getProductsHandler := http.HandlerFunc(getNames)
	http.Handle("/names", getProductsHandler)
	http.ListenAndServe(":8080", nil)
}

func getNames(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names, ok := query["name"]
	if !ok || len(names) == 0 {
		fmt.Println("names not present")
	}
	w.WriteHeader(200)
	b, _ := json.Marshal(names)
	w.Write(b)
}

func doRequest(names ...string) {
	base, err := url.Parse("http://localhost:8080/names")

	if err != nil {
		panic(err)
	}

	params := url.Values{}

	for _, val := range names {
		params.Add("name", val)
	}

	base.RawQuery = params.Encode()
	resp, err := http.Get(base.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var jsonResponse Names
	err = json.Unmarshal(body, &jsonResponse)
	fmt.Println(jsonResponse)
}
