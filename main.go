package main

import (
	"fmt"
	"log"
	"net/http"
)

var states map[string]bool

func main() {

	states = make(map[string]bool)
	http.HandleFunc("/status/", handlerStatus)
	http.HandleFunc("/toggle/", handlerToggle)
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("statusO")

	keys, ok := r.URL.Query()["id"]

	if !ok {
		return
	}

	fmt.Fprintf(w, "%t", states[keys[0]])
}

func handlerToggle(w http.ResponseWriter, r *http.Request) {
	log.Println("toggle")

	keys, ok := r.URL.Query()["id"]
	log.Println(keys)

	if !ok {
		return
	}

	id := keys[0]
	states[id] = !states[id]
	log.Printf("ID %s is now: %t\n", id, states[id])

}
