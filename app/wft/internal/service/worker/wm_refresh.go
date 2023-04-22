package worker

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/yguilai/pipiao-bot/app/wft/internal/domain"
	"io"
	"net/http"
)

const (
	wmAllItemApi = "https://api.warframe.market/v1/items"
	wmMeiliIndex = "pipiao_wm_items"
)

var wmHeader = map[string]string{
	"Language":           "zh-hans",
	"sec-ch-ua-platform": "Windows",
	"Accept":             "application/json",
	"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) " +
		"Chrome/112.0.0.0 Safari/537.36",
	"sec-ch-ua": "\"Chromium\";v=\"112\", \"Google Chrome\";v=\"112\", \"Not:A-Brand\";v=\"99\"",
}

func getAllWmEntry() ([]domain.WarframeMarketItem, error) {
	request, err := http.NewRequest(http.MethodGet, wmAllItemApi, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range wmHeader {
		request.Header.Set(k, v)
	}
	c := &http.Client{}
	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp domain.WmResponse[domain.WmItemResp]
	err = sonic.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Payload == nil {
		return nil, errors.New("data is empty")
	}
	return resp.Payload.Items, nil
}
