package bgta

import (
    "log"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)



func init() {
    r := mux.NewRouter()

    r.HandleFunc("/profile", coolHandler)
    r.HandleFunc("/events", coolHandler)
    r.HandleFunc("/", handler)

    http.Handle("/", r)
    log.Println("Listening...")
    http.ListenAndServe(":3000", nil)
}

func coolHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "BACKEND RULES")
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
