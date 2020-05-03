package main

func main() {
	OAuthResponse := GetAccessToken()
	AuthorizationResponse := Authorization(OAuthResponse.AccessToken)
	Confirmation(OAuthResponse.AccessToken, AuthorizationResponse.Payment.PaymentId)
	Void(OAuthResponse.AccessToken, AuthorizationResponse.Payment.PaymentId)
}