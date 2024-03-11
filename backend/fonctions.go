package backend

import (
	"encoding/json"
	"net/http"
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