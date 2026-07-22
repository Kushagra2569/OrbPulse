package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchType(ctx context.Context, client *http.Client, TypeValue ItemType) (MarketResponse, error) {
	url, _ := buildUrl(TypeValue, RunesOfAldur)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("req err: ", err)
		return MarketResponse{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("resp err: ", err)
		return MarketResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Status:", resp.Status)
		fmt.Println("Body:", string(body))
		return MarketResponse{}, err
	}

	var response MarketResponse

	err = json.NewDecoder(resp.Body).Decode(&response)

	return response, nil
}

func FetchItemDetails(ctx context.Context, client *http.Client, TypeValue ItemType, detailsID string) (ItemDetailResponse, error) {
	url, _ := buildDetailsUrl(TypeValue, RunesOfAldur, detailsID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("req err: ", err)
		return ItemDetailResponse{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("resp err: ", err)
		return ItemDetailResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Status:", resp.Status)
		fmt.Println("Body:", string(body))
		return ItemDetailResponse{}, err
	}

	var response ItemDetailResponse

	err = json.NewDecoder(resp.Body).Decode(&response)

	return response, nil
}
