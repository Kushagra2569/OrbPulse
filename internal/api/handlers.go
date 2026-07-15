package api

import (
	"context"
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

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Body:", string(body))
	} else {
		fmt.Println("Status:", resp.Status)
		fmt.Println("Body:", string(body))
	}

}
