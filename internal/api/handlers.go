package api

import (
	"context"
	"fmt"
	"net/http"
)

func FetchCurrency(ctx context.Context, client *http.Client) {
	fmt.Println("handler called")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, currencyurl, nil)
	if err != nil {
		fmt.Println("req err: ", err)
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println(resp.Body)
	}

}
