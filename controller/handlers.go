package controller

import (
	"groupie/templates"
	"net/http"
)

func Accueilpage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "accueil", nil)
}