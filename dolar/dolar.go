package dolar

import (
	"encoding/json"
	"net/http"
	"time"
)

const url = "https://api.bluelytics.com.ar/v2/latest"

type values struct {
	Buy  float64 `json:"value_buy,omitempty"`
	Sell float64 `json:"value_sell,omitempty"`
	Avg  float64 `json:"value_avg,omitempty"`
}

type CoinValues struct {
	OficialUSD values    `json:"oficial,omitempty"`
	BlueUSD    values    `json:"blue,omitempty"`
	OficialEUR values    `json:"oficial_euro,omitempty"`
	BlueEUR    values    `json:"blue_euro,omitempty"`
	LastUpdate time.Time `json:"last_update,omitempty"`
}

func GetDolar() (*CoinValues, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	var response CoinValues
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, err
}
