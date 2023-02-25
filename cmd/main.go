package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Hello)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()

}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	url := r.URL
	query := url.Query()
	name := query.Get("name")
	surname := query.Get("surname")

	w.Write([]byte("It's not your business"))
	w.Write([]byte("\n" + "I am ..."))

	w.Write([]byte(fmt.Sprintf("\nI am %s %s", name, surname)))
}
