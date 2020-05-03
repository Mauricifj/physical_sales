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

	CheckStatusCode(response)

	if response.StatusCode() == 201 {
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

	fmt.Println("Error on authorizing payment")
	return nil
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
	MerchantOrderId string  `json:"MerchantOrderId"`
	Payment         Payment `json:"Payment"`
}

type Payment struct {
	Amount        int    `json:"Amount"`
	ReceivedDate  string `json:"ReceivedDate"`
	CapturedDate  string `json:"CapturedDate"`
	Status        int    `json:"Status"`
	ReturnMessage string `json:"ReturnMessage"`
	PaymentId     string `json:"PaymentId"`
}

func Confirmation(accessToken string, paymentId string) *ConfirmationResponse {
	configurations := GetConfigurations()

	request := setupConfirmationRequest(accessToken)

	url := configurations.Payment.BaseUrl + strings.Replace(configurations.Payment.Confirmation, "PaymentId", paymentId, -1)

	response, err := request.Put(url)

	if err != nil {
		panic(fmt.Errorf("Fatal error on confirmation request: %s \n", err))
	}

	CheckStatusCode(response)

	if response.StatusCode() == 200 {
		ConfirmationResponse := response.Result().(*ConfirmationResponse)

		fmt.Println("CONFIRMATION STEP")
		fmt.Println(" - PaymentId:", paymentId)
		fmt.Println(" - ConfirmationStatus:", ConfirmationResponse.ConfirmationStatus)
		fmt.Println(" - Status:", ConfirmationResponse.Status)
		fmt.Println(" - ReturnMessage:", ConfirmationResponse.ReturnMessage)
		fmt.Println()

		return ConfirmationResponse
	}
	fmt.Println("Error on confirming payment")
	return nil
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
	ConfirmationStatus int    `json:"ConfirmationStatus"`
	Status             int    `json:"Status"`
	ReturnMessage      string `json:"ReturnMessage"`
}

func Void(accessToken string, paymentId string) *VoidResponse {
	configurations := GetConfigurations()

	request := setupVoidRequest(accessToken)

	url := configurations.Payment.BaseUrl + strings.Replace(configurations.Payment.Void, "PaymentId", paymentId, -1)

	response, err := request.Post(url)

	if err != nil {
		panic(fmt.Errorf("Fatal error on void request: %s \n", err))
	}

	CheckStatusCode(response)

	if response.StatusCode() == 201 {
		VoidResponse := response.Result().(*VoidResponse)

		fmt.Println("VOID STEP")
		fmt.Println(" - PaymentId:", paymentId)
		fmt.Println(" - VoidId:", VoidResponse.VoidId)
		fmt.Println(" - CancellationStatus:", VoidResponse.CancellationStatus)
		fmt.Println(" - Status:", VoidResponse.Status)
		fmt.Println(" - ReturnMessage:", VoidResponse.ReturnMessage)
		fmt.Println()

		return VoidResponse
	}
	fmt.Println("Error on voiding payment")
	return nil
}

func setupVoidRequest(accessToken string) *resty.Request {
	request := resty.New().R()

	request.SetAuthToken(accessToken)
	request.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})
	voidPayload := `{
		"MerchantVoidId": 2019042204,
		"MerchantVoidDate": "2019-04-15T12:00:00Z",
		"Card": {
			"InputMode": "Typed",
			"CardNumber": 5432123454321234
		}
	}`
	request.SetBody(voidPayload)
	request.SetResult(VoidResponse{})

	return request
}

type VoidResponse struct {
	VoidId             string `json:"VoidId"`
	CancellationStatus int    `json:"CancellationStatus"`
	Status             int    `json:"Status"`
	ReturnMessage      string `json:"ReturnMessage"`
}

func UndoVoid(accessToken string, paymentId string, voidId string) *UndoVoidResponse {
	configurations := GetConfigurations()

	request := setupUndoVoidRequest(accessToken)

	paymentIdReplaced := strings.Replace(configurations.Payment.UndoVoid, "PaymentId", paymentId, -1)
	voidIdReplaced := strings.Replace(paymentIdReplaced, "VoidId", voidId, -1)

	url := configurations.Payment.BaseUrl + voidIdReplaced

	response, err := request.Delete(url)

	if err != nil {
		panic(fmt.Errorf("Fatal error on undo void request: %s \n", err))
	}

	CheckStatusCode(response)

	if response.StatusCode() == 200 {
		UndoVoidResponse := response.Result().(*UndoVoidResponse)

		fmt.Println("UNDO VOID STEP")
		fmt.Println(" - PaymentId:", paymentId)
		fmt.Println(" - VoidId:", voidId)
		fmt.Println(" - CancellationStatus:", UndoVoidResponse.CancellationStatus)
		fmt.Println(" - Status:", UndoVoidResponse.Status)
		fmt.Println(" - ReturnMessage:", UndoVoidResponse.ReturnMessage)
		fmt.Println()

		return UndoVoidResponse
	}
	fmt.Println("Error on undoing void of payment")
	return nil
}

func setupUndoVoidRequest(accessToken string) *resty.Request {
	request := resty.New().R()

	request.SetAuthToken(accessToken)
	request.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})
	request.SetResult(UndoVoidResponse{})

	return request
}

type UndoVoidResponse struct {
	CancellationStatus int    `json:"CancellationStatus"`
	Status             int    `json:"Status"`
	ReturnMessage      string `json:"ReturnMessage"`
}

func CheckStatusCode(response *resty.Response) {
	switch response.StatusCode() {
	case 400:
		panic(fmt.Errorf("Error with contract: %s \n", response.Status()))

	case 401:
		panic(fmt.Errorf("Error with credentials: %s \n", response.Status()))

	case 403:
		panic(fmt.Errorf("Operation not allowed: %s \n", response.Status()))

	case 404:
		panic(fmt.Errorf("Not found or undone: %s \n", response.Status()))
	}
}
