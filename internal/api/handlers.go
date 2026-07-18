package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchCurrency(ctx context.Context, client *http.Client) {
	fmt.Println("handler called")
	currencyurl, _ := buildUrl(Currency, RunesOfAldur)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, currencyurl, nil)
	if err != nil {
		fmt.Println("req err: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("resp err: ", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Status:", resp.Status)
		fmt.Println("Body:", string(body))
		return
	}

	var response Response

	err = json.NewDecoder(resp.Body).Decode(&response)

	fmt.Println(response.Core.Rates)

}
