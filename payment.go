package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
)

func Authorization(accessToken string) *AuthorizationResponse {
	configurations := GetConfigurations()

	request := setupAuthorizationRequest(accessToken)

	url := configurations.Payment.BaseUrl + configurations.Payment.Authorization

	response, err := request.Post(url)

	if err != nil {
		panic(fmt.Errorf("Fatal error on authorization request: %s \n", err))
	}

	AuthorizationResponse := response.Result().(*AuthorizationResponse)

	fmt.Println("AUTHORIZATION STEP")
	fmt.Println(" - MerchantOrderId:", AuthorizationResponse.MerchantOrderId)
	fmt.Println(" - Amount:", AuthorizationResponse.Payment.Amount)
	fmt.Println(" - ReceivedDate:", AuthorizationResponse.Payment.ReceivedDate)
	fmt.Println(" - CapturedDate:", AuthorizationResponse.Payment.CapturedDate)
	fmt.Println(" - Status:", AuthorizationResponse.Payment.Status)
	fmt.Println(" - ReturnMessage:", AuthorizationResponse.Payment.ReturnMessage)
	fmt.Println(" - PaymentId:", AuthorizationResponse.Payment.PaymentId)
	fmt.Println()

	return AuthorizationResponse
}

func setupAuthorizationRequest(accessToken string) *resty.Request {
	request := resty.New().R()

	request.SetAuthToken(accessToken)
	request.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})

	paymentPayload := `{
		"MerchantOrderId": "1587997030607",
			"Payment": {
			"Type": "PhysicalCreditCard",
				"SoftDescriptor": "Desafio GO 2",
				"PaymentDateTime": "2020-01-08T11:00:00",
				"Amount": 100,
				"Installments": 1,
				"Interest": "ByMerchant",
				"Capture": true,
				"ProductId": 1,
				"CreditCard": {
				"CardNumber": "5432123454321234",
					"ExpirationDate": "12/2021",
					"SecurityCodeStatus": "Collected",
					"SecurityCode": "123",
					"BrandId": 1,
					"IssuerId": 401,
					"InputMode": "Typed",
					"AuthenticationMethod": "NoPassword",
					"TruncateCardNumberWhenPrinting": true
			},
			"PinPadInformation": {
				"PhysicalCharacteristics": "PinPadWithChipReaderWithoutSamAndContactless",
					"ReturnDataInfo": "00",
					"SerialNumber": "0820471929",
					"TerminalId": "42004558"
			}
		}
	}`
	request.SetBody(paymentPayload)

	request.SetResult(AuthorizationResponse{})

	return request
}

type AuthorizationResponse struct {
	MerchantOrderId string `json:"MerchantOrderId"`
	Payment Payment `json:"Payment"`
}

type Payment struct {
	Amount int `json:"Amount"`
	ReceivedDate string `json:"ReceivedDate"`
	CapturedDate string `json:"CapturedDate"`
	Status int `json:"Status"`
	ReturnMessage string `json:"ReturnMessage"`
	PaymentId string `json:"PaymentId"`
}

func Confirmation(accessToken string, paymentId string) *ConfirmationResponse {
	configurations := GetConfigurations()

	request := setupConfirmationRequest(accessToken)

	url := configurations.Payment.BaseUrl + strings.Replace(configurations.Payment.Confirmation, "PaymentId", paymentId, -1)

	response, err := request.Put(url)

	if err != nil {
		panic(fmt.Errorf("Fatal error on confirmation request: %s \n", err))
	}

	ConfirmationResponse := response.Result().(*ConfirmationResponse)

	fmt.Println("CONFIRMATION STEP")
	fmt.Println(" - ConfirmationStatus:", ConfirmationResponse.ConfirmationStatus)
	fmt.Println(" - Status:", ConfirmationResponse.Status)
	fmt.Println(" - ReturnMessage:", ConfirmationResponse.ReturnMessage)
	fmt.Println()

	return ConfirmationResponse
}

func setupConfirmationRequest(accessToken string) *resty.Request {
	request := resty.New().R()

	request.SetAuthToken(accessToken)
	request.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})
	request.SetBody("{}")
	request.SetResult(ConfirmationResponse{})

	return request
}

type ConfirmationResponse struct {
	ConfirmationStatus int `json:"ConfirmationStatus"`
	Status int `json:"Status"`
	ReturnMessage string `json:"ReturnMessage"`
}