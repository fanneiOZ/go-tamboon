package v20190529

import (
	"encoding/json"
	"fmt"
	"go-tamboon/internal/infrastructure/external/omise"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type TokenAPI struct {
	baseURL   string
	publicKey string
	http      *omise.HttpClient
}

func NewTokenAPI(cfg omise.ApiConfig) *TokenAPI {
	return &TokenAPI{
		baseURL:   cfg.Endpoints.Vault,
		publicKey: cfg.PublicKey,
		http:      omise.NewHttpClient(),
	}
}

func (api *TokenAPI) CreateToken(input TokenRequest) (*TokenResponse, error) {
	form := url.Values{}
	inputCardValues := reflect.ValueOf(input.Card)
	inputCardKeys := reflect.TypeOf(input.Card)
	for i := 0; i < inputCardValues.NumField(); i++ {
		field := inputCardKeys.Field(i)
		inputValue := inputCardValues.Field(i)
		tag := field.Tag.Get("json")
		key := strings.Split(tag, ",")[0]
		if inputValue.Kind() == reflect.Ptr && inputValue.IsNil() {
			continue
		}

		value := ""
		if inputValue.Kind() == reflect.Ptr {
			value = fmt.Sprintf("%v", inputValue.Elem().Interface())
		} else {
			value = fmt.Sprintf("%v", inputValue.Interface())
		}

		form.Set(fmt.Sprintf("card[%s]", key), value)
	}

	form.Set("object", "card")

	req, _ := http.NewRequest("POST", api.baseURL+"/tokens", strings.NewReader(form.Encode()))
	req.SetBasicAuth(api.publicKey, "")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	body, err := api.http.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var result TokenResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
