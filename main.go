package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	url = "http://applog.apps.axonerp.com/geotech/logs?auth=uJFb6j0flWctLkBZhgjPPIj2GeqzUpA"
)

func logs(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-type", "application/json")
	io.WriteString(w, string(body))
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	fmt.Println("Welcome to Logs")
	server := http.Server{
		Addr:    ":9000",
		Handler: &myHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = logs
	server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "No handler found on server: "+r.URL.String())
}
