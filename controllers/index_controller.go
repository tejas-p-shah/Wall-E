package controllers

import (
	"html/template"
	"net/http"
)

type IndexPage struct {
	Title string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	data := IndexPage{Title: "Wall-E Home"}

	t := template.Must(template.ParseFiles("views/templates/index.gohtml"))
	t.Execute(w, data)
}
