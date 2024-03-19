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

func NewsPage(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://api.mozambiquehe.re/news?auth=96b5ef5aff1825ee880530ccf48efd82"
	news, err := backend.FetchNewsData(apiURL)
	if err != nil {
		fmt.Println("Error fetching news data:", err)
		return
	}
	templates.Temp.ExecuteTemplate(w, "news", news)
	
}

func InfoPage(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://api.mozambiquehe.re/predator?auth=96b5ef5aff1825ee880530ccf48efd82"
	info, err := backend.FetchInfoData(apiURL)
	if err != nil {
		fmt.Println("Error fetching info data:", err)
		return
	}
	templates.Temp.ExecuteTemplate(w, "info", info)
}

func RotaPage(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://api.mozambiquehe.re/maprotation?auth=96b5ef5aff1825ee880530ccf48efd82"
	rota, err := backend.FetchRotaData(apiURL)
	if err != nil {
		fmt.Println("Error fetching rota data:", err)
		return
	}
	templates.Temp.ExecuteTemplate(w, "rota", rota)
}
