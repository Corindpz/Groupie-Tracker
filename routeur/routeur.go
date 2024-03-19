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
	http.HandleFunc("/shop", controller.ShopPage)
	http.HandleFunc("/news", controller.NewsPage)
	http.HandleFunc("/info", controller.InfoPage)
	http.HandleFunc("/rota", controller.RotaPage)
	log.Println("Serveur lancé")
	http.ListenAndServe(":8080", nil)
	log.Fatal("Serveur arrêté")
}