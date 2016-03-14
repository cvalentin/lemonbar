package main

import (
    "log"
    "net/http"
)

func main() {
    http.Handle("/", getRoutes())

    log.Println("Listening...")
    http.ListenAndServe(":3000", nil)
}
