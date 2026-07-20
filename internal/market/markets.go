package market

import (
	"github.com/Kushagra2569/OrbPulse/internal/api"
)

func (m *Markets) Set(typevalue api.ItemType, market *Market) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[typevalue] = market
}

func (m *Markets) GetMarket(typevalue api.ItemType) (*Market, bool) {
	marketValue, exists := m.data[typevalue]
	return marketValue, exists
}

func (m *Markets) GetItem(typevalue api.ItemType, id string) (api.Item, bool) {
	marketValue, exists := m.data[typevalue]

	if !exists {
		return api.Item{}, exists
	}

	return marketValue.Items[id], exists
}
