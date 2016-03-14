package main

import (
    "log"
    "fmt"
    "net/http"
    "io/ioutil"
    "github.com/gorilla/mux"
)

func getRoutes() *mux.Router {
    r := mux.NewRouter()

    r.HandleFunc("/profile", profilerHandler)
    r.HandleFunc("/events", coolHandler)

    staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
    r.PathPrefix("/static").Handler(staticHandler)

    r.HandleFunc("/", rootHandler)
    return r
}

func profilerHandler(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadFile("views/profile.html")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Fprintf(w, "%s", string(body))
}

func coolHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "BACKEND RULES")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
