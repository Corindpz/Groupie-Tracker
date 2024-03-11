package controller

import (
	"groupie/templates"
	"net/http"
	"fmt"
	"groupie/backend"
)

func Accueilpage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "accueil", nil)
}

func ShopPage(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://api.mozambiquehe.re/store?auth=96b5ef5aff1825ee880530ccf48efd82"
	bundles, err := backend.FetchStoreData(apiURL)
	if err != nil {
		fmt.Println("Error fetching store data:", err)
		return
	}

	templates.Temp.ExecuteTemplate(w, "shop", bundles)
}
