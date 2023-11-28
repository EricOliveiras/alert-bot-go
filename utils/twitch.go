package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ericoliveiras/alert-bot-go/response"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func TwitchGenerateToken() (string, error) {

	data := map[string]string{
		"client_id":     cfg.Twitch.ClientID,
		"client_secret": cfg.Twitch.ClientSecret,
		"grant_type":    "client_credentials",
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(cfg.Twitch.TokenURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP error: %d - %s", resp.StatusCode, string(body))
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

func GetStream(login string) (response.StreamResponse, error) {
	token, err := TwitchGenerateToken()
	if err != nil {
		return response.StreamResponse{}, fmt.Errorf("error to generate token: %s", err)
	}

	req, err := http.NewRequest("GET", cfg.Twitch.SearchStramURL, nil)
	if err != nil {
		return response.StreamResponse{}, fmt.Errorf("error to create request: %s", err)
	}

	q := req.URL.Query()
	q.Add("query", strings.ToLower(login))
	q.Add("first", "1")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Client-Id", cfg.Twitch.ClientID)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response.StreamResponse{}, fmt.Errorf("error to send request: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response.StreamResponse{}, fmt.Errorf("error reading the answer: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return response.StreamResponse{}, fmt.Errorf("HTTP error: %d - %s", resp.StatusCode, string(body))
	}

	var streamData response.StreamResponse
	err = json.Unmarshal(body, &streamData)
	if err != nil {
		return response.StreamResponse{}, fmt.Errorf("error when parsing JSON: %s", err)
	}

	return streamData, nil
}
