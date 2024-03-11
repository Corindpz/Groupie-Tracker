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
<<<<<<< HEAD
	http.HandleFunc("/shop", controller.ShopPage)
=======
>>>>>>> 5009e5137a04e6fbc208e7da539cf3e916a6bf7c

	
	log.Println("Serveur lancé")
	http.ListenAndServe(":8080", nil)
	log.Fatal("Serveur arrêté")
}