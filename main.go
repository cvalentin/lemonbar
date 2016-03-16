package main

import (
    "log"
    "net/http"
)

func getPort() string {
        var port = os.Getenv("PORT")
        // Set a default port if there is nothing in the environment
        if port == "" {
                port = "3000"
                fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
        }
        return ":" + port
}

func main() {
    http.Handle("/", getRoutes())

    log.Println("Listening...")
    http.ListenAndServe(getPort(), nil)
}
