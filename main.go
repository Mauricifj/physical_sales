package main

func main() {
	OAuthResponse := GetAccessToken()
	Authorization(OAuthResponse.AccessToken)
}