package service

import (
	"net/http"
	"text/template"

)

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username :=  r.Form["username"][0]
	password :=  r.Form["password"][0]

	t := template.Must(template.ParseFiles("templates/info.html"))
	t.Execute(w, map[string]string{
		"username": username,
		"password": password,
	})

}