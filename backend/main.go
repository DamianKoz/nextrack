package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":8000"

	mux := http.NewServeMux()
	mux.HandleFunc("/events", getAllEvents)

	http.HandleFunc("/events", getAllEvents)
	fmt.Printf("Listening on Port %s\n", port)
	http.ListenAndServe(port, nil)
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("alle Events"))
}
