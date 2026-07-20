package main

import (
	"context"
	"fmt"
	"github.com/Kushagra2569/OrbPulse/internal/api"
	"github.com/Kushagra2569/OrbPulse/internal/market"
	"net/http"
	"time"
)

func main() {
	fmt.Println("start")

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	ctx := context.Background()

	var markets = market.NewMarkets()

	for _, TypeValue := range api.AllItemTypes {
		response, err := api.FetchType(ctx, client, TypeValue)
		if err != nil {
			fmt.Println("fetch failed for type ", TypeValue)
			continue
		}
		marketValue := market.FromResponse(response, TypeValue)
		markets.Set(TypeValue, marketValue)
	}

	for _, typeValue := range api.AllItemTypes {
		market, _ := markets.GetMarket(typeValue)
		fmt.Println((*market).Core)
	}

}
