package controller

import (
	"groupie/templates"
	"net/http"
<<<<<<< HEAD
	"fmt"
	"groupie/backend"
=======
>>>>>>> 5009e5137a04e6fbc208e7da539cf3e916a6bf7c
)

func Accueilpage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "accueil", nil)
<<<<<<< HEAD
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
=======
}
>>>>>>> 5009e5137a04e6fbc208e7da539cf3e916a6bf7c
