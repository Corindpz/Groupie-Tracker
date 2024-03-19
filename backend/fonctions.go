
package backend

import (
	"encoding/json"
	"net/http"
	"fmt"
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