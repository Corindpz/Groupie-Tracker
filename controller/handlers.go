package controller

import (
	"groupie/templates"
	"net/http"
	"fmt"
	"groupie/backend"
	"strconv"
)

func Accueilpage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "accueil", nil)

}

func TrackerPage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "tracker", nil)
}

func SearchTrackerPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pseudo := r.FormValue("pseudo")
	platform := r.FormValue("platform")

	fmt.Println("Pseudo:", pseudo)
	fmt.Println("Platform:", platform)
	apiURL := "https://api.mozambiquehe.re/bridge?platform=" + platform + "&player=" + pseudo + "&auth=96b5ef5aff1825ee880530ccf48efd82"

	fmt.Println("API URL:", apiURL)
	fmt.Println("iciciciciciciiciiciciciicicic")

	tracker, err := backend.FetchTrackerData(apiURL)
	if err != nil {
		fmt.Println("Error fetching tracker data:", err)
		templates.Temp.ExecuteTemplate(w, "error", nil)
		return
	}
	fmt.Println(tracker.Total)

	fmt.Println("Tracker:", tracker)
	templates.Temp.ExecuteTemplate(w, "trackerstats", tracker)

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
	page := r.FormValue("page")
	currentPage, errConv := strconv.Atoi(page)
	if(errConv != nil || page == "" || currentPage < 1) {
		currentPage = 1
	}
	fmt.Println(len(news), currentPage*4)
	if (len(news) < currentPage*4) {
		fmt.Println("Error: page number too high")
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}

	news = news[(currentPage*4)-4:currentPage*4]
	var dataPage backend.NewsData = backend.NewsData{News: news, PageNext:fmt.Sprintf("/news?page=%v",currentPage+1), PagePrev: fmt.Sprintf("/news?page=%v",currentPage-1)}
	templates.Temp.ExecuteTemplate(w, "news", dataPage)
	
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

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "error", nil)

}