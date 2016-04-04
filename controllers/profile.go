package controllers

import (
    "log"
    "net/http"
    "html/template"
)

const TemplatePath string = "views/"

func GetProfileIndex(w http.ResponseWriter, r *http.Request){

    page, err := template.New("profile").ParseFiles(TemplatePath + "profile.tmpl")

    if err!= nil {
        log.Fatal(err)
    }

    page, err = page.ParseGlob(TemplatePath + "common/*")

    if err!= nil {
        log.Fatal(err)
    }

    err = page.ExecuteTemplate(w,"profile", nil)

    if err != nil {
        log.Fatal(err)
    }
}

func onLoad() string {
    return "Ash Douglas Hill"
}
