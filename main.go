package main

func main() {
	OAuthResponse := GetAccessToken()
	AuthorizationResponse := Authorization(OAuthResponse.AccessToken)
	Confirmation(OAuthResponse.AccessToken, AuthorizationResponse.Payment.PaymentId)
	VoidResponse := Void(OAuthResponse.AccessToken, AuthorizationResponse.Payment.PaymentId)
	UndoVoid(OAuthResponse.AccessToken, AuthorizationResponse.Payment.PaymentId, VoidResponse.VoidId)
}