package market

import (
	"sync"

	"github.com/Kushagra2569/OrbPulse/internal/api"
)

type Market struct {
	Type api.ItemType

	Core  api.Core
	Lines []api.MarketLine

	Items map[string]api.Item
}

type Markets struct {
	mu sync.RWMutex

	data map[api.ItemType]*Market
}

func NewMarkets() *Markets {
	return &Markets{
		data: make(map[api.ItemType]*Market),
	}
}

func FromResponse(response api.MarketResponse, TypeValue api.ItemType) *Market {

	itemsMap := make(map[string]api.Item)

	for _, item := range response.Items {
		itemsMap[item.ID] = item
	}

	return &Market{
		Type:  TypeValue,
		Core:  response.Core,
		Lines: response.Lines,
		Items: itemsMap,
	}
}
