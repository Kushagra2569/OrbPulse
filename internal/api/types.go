package api

type ItemType string

const (
	Currency           ItemType = "Currency"
	Fragments          ItemType = "Fragments"
	Abyss              ItemType = "Abyssal Bones"
	UncutGems          ItemType = "Uncut Gems"
	LineageSupportGems ItemType = "Lineage Gems"
	Essences           ItemType = "Essences"
	SoulCores          ItemType = "Soul Cores"
	Idols              ItemType = "Idols"
	Runes              ItemType = "Runes"
	Ritual             ItemType = "Omens"
	Expedition         ItemType = "Expedition"
	Delirium           ItemType = "Liquid Emotions"
	Breach             ItemType = "Catalysts"
	Verisium           ItemType = "Verisium"
)

var AllItemTypes = []ItemType{
	Currency,
	Fragments,
	Abyss,
	UncutGems,
	LineageSupportGems,
	Essences,
	SoulCores,
	Idols,
	Runes,
	Ritual,
	Expedition,
	Delirium,
	Breach,
	Verisium,
}

type League string

const (
	RunesOfAldur League = "Runes of Aldur"
	Hardcore     League = "Hardcore"
	Standard     League = "Standard"
)
