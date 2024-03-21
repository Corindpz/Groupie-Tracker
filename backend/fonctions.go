
package backend

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
	"errors"
)

func FetchStoreData(apiURL string) ([]Bundle, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var bundles []Bundle
	err = json.NewDecoder(resp.Body).Decode(&bundles)
	if err != nil {
		return nil, err
	}

	return bundles, nil
}

func FetchTrackerData(apiURL string) (Tracker, error) {
    fmt.Println("API URL2:", apiURL)

    resp, err := http.Get(apiURL)
    if err != nil {
		fmt.Println("Erreur lors de la requête GET:", err)
        return Tracker{}, fmt.Errorf("erreur lors de la requête GET: %v", err)
    }

	if  resp.StatusCode == 400 {
		return Tracker{}, errors.New("Joueur non trouvé")
	}

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
		fmt.Println("Erreur: statut de la réponse", resp.StatusCode)
        return Tracker{}, fmt.Errorf("erreur: statut de la réponse %d", resp.StatusCode)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		fmt.Println("Erreur lors de la lecture du corps de la réponse:", err)
        return Tracker{}, fmt.Errorf("erreur lors de la lecture du corps de la réponse: %v", err)
    }

    if len(body) == 0 {
		fmt.Println("Aucune donnée dans la réponse")
        return Tracker{}, errors.New("aucune donnée dans la réponse")
    }

    var tracker Tracker
    err = json.Unmarshal(body, &tracker)
    if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
        return Tracker{}, fmt.Errorf("erreur lors du décodage JSON: %v", err)
    }

    return tracker, nil
}


func FetchNewsData(apiURL string) ([]News, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var news []News
	err = json.NewDecoder(resp.Body).Decode(&news)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func FetchInfoData(apiURL string) (RotaData, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return RotaData{}, err
	}
	defer resp.Body.Close()

	var rotaData RotaData

	if err := json.NewDecoder(resp.Body).Decode(&rotaData); err != nil {
		return RotaData{}, fmt.Errorf("erreur lors de la désérialisation JSON : %w", err)
	}

	fmt.Println(rotaData)
	return rotaData, nil
}

func FetchRotaData(apiURL string) (MapRotation, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return MapRotation{}, err
	}
	defer resp.Body.Close()

	var rota MapRotation
	err = json.NewDecoder(resp.Body).Decode(&rota)
	if err != nil {
		return MapRotation{}, err
	}

	return rota, err
}