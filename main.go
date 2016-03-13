package main

import (
    "log"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "HELLO WORLD", r.URL.Path[1:])
}

func coolHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "BACKEND RULES")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", handler)
    r.HandleFunc("/profile", coolHandler)
    r.HandleFunc("/events", coolHandler)

    http.Handle("/", r)
    log.Println("Listening...")
    http.ListenAndServe(":3000", nil)
}
