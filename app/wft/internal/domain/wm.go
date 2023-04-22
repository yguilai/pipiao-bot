package domain

type (
	WmResponse[T any] struct {
		Payload *T `json:"payload"`
	}

	WmItemResp struct {
		Items []WarframeMarketItem `json:"items"`
	}

	WarframeMarketItem struct {
		Id       string `json:"id,omitempty"`
		Thumb    string `json:"thumb,omitempty"`
		UrlName  string `json:"url_name,omitempty"`
		ItemName string `json:"item_name,omitempty"`
		Vaulted  bool   `json:"vaulted,omitempty"`
	}
)
