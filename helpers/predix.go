package helpers

import (
	"log"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

// GetPredixClientConfig ...
func GetPredixClientConfig(config *UserConfig) *cfclient.Config {
	return &cfclient.Config{
		ApiAddress: config.Predix.API,
		Username:   config.Predix.Username,
		Password:   config.Predix.Password,
	}
}

// PredixLogin ...
func PredixLogin(clientConfig *cfclient.Config) *cfclient.Client {
	log.Println("-> Logging to Predix.io ...")
	client, err := cfclient.NewClient(clientConfig)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("* LOGGED IN as %s", clientConfig.Username)
	return client
}
