package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Credentials struct {
	Username string `yaml:username`
	Password string `yaml:password`
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