package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v3"
	"os"
)

func GetAccessToken() *OAuthResponse {
	configurations := GetConfigurations()

	request := setupOAuthRequest()

	url := configurations.Authentication.BaseUrl + configurations.Authentication.OAuth

	response, err := request.Post(url)

	if err != nil {
		panic(fmt.Errorf("Fatal error on oauth request: %s \n", err))
	}

	OAuthResponse := response.Result().(*OAuthResponse)

	fmt.Println("AUTHENTICATION STEP")
	fmt.Println(" - access_token:", OAuthResponse.AccessToken[0:60], "...")
	fmt.Println(" - token_type:", OAuthResponse.TokenType)
	fmt.Println(" - expires_in:", OAuthResponse.ExpiresIn)
	fmt.Println()

	return OAuthResponse
}

func setupOAuthRequest() *resty.Request {
	credentials := GetCredentials()

	request := resty.New().R()

	request.SetBasicAuth(credentials.Username, credentials.Password)

	request.SetFormData(map[string]string{
		"grant_type": "client_credentials",
	})

	request.SetResult(OAuthResponse{})
	return request
}

func GetCredentials() *Credentials {
	credentials := &Credentials{}
	path := "./credentials.yml"

	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("Fatal error while opening credentials: %s \n", err))
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&credentials); err != nil {
		panic(fmt.Errorf("Fatal error while decoding credentials: %s \n", err))
	}

	return credentials
}

type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type Credentials struct {
	Username string `yaml:username`
	Password string `yaml:password`
}