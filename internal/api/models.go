package api

type CoreCurrency struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	Category  string `json:"category"`
	DetailsID string `json:"detailsId"`
}

type Core struct {
	Items     []CoreCurrency     `json:"items"`
	Rates     map[string]float64 `json:"rates"`
	Primary   string             `json:"primary"`
	Secondary string             `json:"secondary"`
}

type Sparkline struct {
	TotalChange float64    `json:"totalChange"`
	Data        []*float64 `json:"data"`
}

type MarketLine struct {
	ID                 string  `json:"id"`
	PrimaryValue       float64 `json:"primaryValue"`
	VolumePrimaryValue float64 `json:"volumePrimaryValue"`

	MaxVolumeCurrency string  `json:"maxVolumeCurrency"`
	MaxVolumeRate     float64 `json:"maxVolumeRate"`

	Sparkline Sparkline `json:"sparkline"`
}

type Item struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	Category  string `json:"category"`
	DetailsID string `json:"detailsId"`
}

type MarketResponse struct {
	Core  Core         `json:"core"`
	Lines []MarketLine `json:"lines"`
	Items []Item       `json:"items"`
}
