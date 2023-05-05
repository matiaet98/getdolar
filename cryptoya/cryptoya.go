package cryptoya

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const url = "https://criptoya.com/api/"

type Values struct {
	Ask      float64 `json:"ask,omitempty"`
	TotalAsk float64 `json:"totalAsk,omitempty"`
	Bid      float64 `json:"bid,omitempty"`
	TotalBid float64 `json:"totalBid,omitempty"`
	Time     int64   `json:"time,omitempty"`
}

func GetCrypto(exchange string, coin string) (*Values, error) {
	client := &http.Client{}
	uri := fmt.Sprintf("%s/%s/%s/ars/1.0", url, exchange, coin)
	req, err := http.NewRequest("GET", uri, http.NoBody)
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
	var response Values
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, err
}
