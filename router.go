package main

import (
    "log"
    "fmt"
    "net/http"
    "io/ioutil"
    "github.com/gorilla/mux"
    "github.com/cvalentin/lemonbar/controllers"
)

func getRoutes() *mux.Router {
    r := mux.NewRouter()

    r.HandleFunc("/group", staticPageHandler)
    r.HandleFunc("/profile", controllers.GetProfileIndex)
    r.HandleFunc("/index", staticPageHandler)
    r.HandleFunc("/library", staticPageHandler)
    r.HandleFunc("/event", staticPageHandler)

    staticAssetsHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
    r.PathPrefix("/static").Handler(staticAssetsHandler)

    r.HandleFunc("/", rootRedirect)

    return r
}

func staticPageHandler(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadFile("views" + r.URL.String() + ".html")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Fprintf(w, "%s", string(body))
}

func rootRedirect(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/index", 301)
}
