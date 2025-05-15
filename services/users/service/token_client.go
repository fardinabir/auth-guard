package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fardinabir/auth-guard/model"
	"github.com/spf13/viper"
)

type TokenClient struct {
	BaseURL string
}

func NewTokenClient() *TokenClient {
	return &TokenClient{
		BaseURL: viper.GetString("services.warden.host") + ":" + viper.GetString("services.warden.port"),
	}
}

func (c *TokenClient) GenerateToken(username string) (model.Token, error) {
	reqBody, err := json.Marshal(map[string]string{
		"username": username,
	})
	if err != nil {
		return model.Token{}, err
	}

	resp, err := http.Post(c.BaseURL+"/tokens", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return model.Token{}, err
	}
	defer resp.Body.Close()

	var token model.Token
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return model.Token{}, err
	}
	return token, nil
}

func (c *TokenClient) ValidateToken(token string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/tokens/validate", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token validation failed")
	}

	var claims map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&claims); err != nil {
		return nil, err
	}
	return claims, nil
}
