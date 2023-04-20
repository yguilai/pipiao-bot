package domain

type WarframeMarketItem struct {
	Id       string `json:"id,omitempty"`
	Thumb    string `json:"thumb,omitempty"`
	UrlName  string `json:"urlName,omitempty"`
	ItemName string `json:"itemName,omitempty"`
	Vaulted  bool   `json:"vaulted,omitempty"`
}
