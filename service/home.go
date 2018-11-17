package service

import (
	"net/http"
	"github.com/unrolled/render"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request)  {

	t := template.Must(template.ParseFiles("templates/login.html"))
	t.Execute(w, nil)
}

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			WELCOME      string `json:"WELCOME"`
			
		}{WELCOME: "Hello"})
}
}
