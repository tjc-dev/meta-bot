package jup

import (
	"encoding/json"
	"net/http"
)

type Price struct {
	Data struct {
		Meta struct {
			ID            string  `json:"id"`
			MintSymbol    string  `json:"mintSymbol"`
			VsToken       string  `json:"vsToken"`
			VsTokenSymbol string  `json:"vsTokenSymbol"`
			Price         float64 `json:"price"`
		} `json:"META"`
	} `json:"data"`
	TimeTaken float64 `json:"timeTaken"`
}

func GetMETAPrice() float64 {
	client := &http.Client{}
	url := "https://price.jup.ag/v4/price?ids=META&vsToken=USDC"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0
	}
	req.Header.Set("user-agent", "Meta BOT Price Fetcher")
	r, err := client.Do(req)
	if err != nil {
		return 0
	}
	defer r.Body.Close()

	var priceData Price

	err = json.NewDecoder(r.Body).Decode(&priceData)

	if err != nil {
		return 0
	}

	return priceData.Data.Meta.Price
}
