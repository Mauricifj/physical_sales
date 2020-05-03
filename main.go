package main

import "fmt"

func main() {
	oauthResponse := GetAccessToken()
	fmt.Println(oauthResponse.AccessToken)
}