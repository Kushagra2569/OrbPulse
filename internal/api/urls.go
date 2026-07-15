package api

import (
	"net/url"
)

const (
	baseURL = "https://poe.ninja/poe2/api/economy/exchange/current"
)

func buildUrl(item ItemType, league League) (string, error) {

	u, _ := url.Parse(baseURL + "/overview")

	q := u.Query()
	q.Set("league", string(league))
	q.Set("type", string(item))

	u.RawQuery = q.Encode()

	return u.String(), nil
}
