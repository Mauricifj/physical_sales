package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Configuration struct {
	Authentication struct{
		BaseUrl string `yaml:baseurl`
		OAuth   string `yaml:oauth`
	}`yaml:authentication`
}

func GetConfigurations() *Configuration {
	configuration := &Configuration{}
	path := "./configuration.yml"

	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("Fatal err while opening configurations: %s \n", err))
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&configuration); err != nil {
		panic(fmt.Errorf("Fatal err while decoding configurations: %s \n", err))
	}
	return configuration
}