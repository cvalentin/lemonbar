package controllers

import (
    "log"
    "fmt"
    "net/http"
    "io/ioutil"
    "path/filepath"
    "os"
)

func GetProfileIndex(w http.ResponseWriter, r *http.Request){
    dir, _ := os.Getwd()
    indexHTML := filepath.Join(dir, "views/profile.html")
    body, err := ioutil.ReadFile(indexHTML)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Fprintf(w, "%s", string(body))
}

func onLoad() string {
    return "Ash Douglas Hill"
}
