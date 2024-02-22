package routeur

import (
	"groupie/controller"
	"log"
	"net/http"
)

func Initserv() {
	css:= http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	http.HandleFunc("/accueil", controller.Accueilpage)

	
	log.Println("Serveur lancé")
	http.ListenAndServe(":8080", nil)
	log.Fatal("Serveur arrêté")
}