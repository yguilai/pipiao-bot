package types

type WarframeItem struct {
	UniqueName string
	Name       string
}

type WarframeItemI18n struct {
	Name        string `json:"name"`
	SystemName  string `json:"systemName"`
	Description string `json:"description"`
	Trigger     string `json:"trigger"`
	LevelStats  []struct {
		Stats []string `json:"stats"`
	} `json:"levelStats"`
}
