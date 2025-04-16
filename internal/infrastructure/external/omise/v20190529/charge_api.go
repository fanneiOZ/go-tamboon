package v20190529

import (
	"encoding/json"
	"fmt"
	"go-tamboon/internal/infrastructure/external/omise"
	"net/http"
	"net/url"
	"strings"
)

type ChargeAPI struct {
	baseURL   string
	secretKey string
	http      *omise.HttpClient
}

func NewChargeAPI(cfg omise.ApiConfig) *ChargeAPI {
	return &ChargeAPI{
		baseURL:   cfg.Endpoints.Api,
		secretKey: cfg.SecretKey,
		http:      omise.NewHttpClient(),
	}
}

func (api *ChargeAPI) CreateCharge(input ChargeRequest) (*ChargeResponse, error) {
	form := url.Values{}
	form.Set("amount", fmt.Sprintf("%d", input.Amount))
	form.Set("currency", input.Currency)
	form.Set("card", input.Card)
	form.Set("capture", "true")

	req, _ := http.NewRequest("POST", api.baseURL+"/charges", strings.NewReader(form.Encode()))
	req.SetBasicAuth(api.secretKey, "")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	body, err := api.http.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var result ChargeResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
