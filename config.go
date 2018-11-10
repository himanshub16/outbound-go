package main

// Credits for this file: https://www.thepolyglotdeveloper.com/2017/04/load-json-configuration-file-golang-application/

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration describes the fields of JSON configuration
type Configuration struct {
	DBTYPE      string `json:"DB_TYPE"`
	DatabaseURL string `json:"DATABASE_URL"`
	LinksColl   string `json:"LINKS_COLL,omitempty"`
	CounterColl string `json:"COUNTER_COLL,omitempty"`

	Port           string `json:"PORT"`
	RedirectMethod string `json:"REDIRECT_METHOD"`
	AuthToken      string `json:"AUTH_TOKEN"`
}

// ReadConfig reads config from config file
func ReadConfig() *Configuration {
	var (
		config         Configuration
		configFilename string
	)

	configFilename = os.Getenv("CONFIG_FILE")
	//configFilename = ".env.postgresql.json"
	//configFilename = ".env.mongodb.json"
	if len(configFilename) == 0 {

		log.Println("Config file not provided. Reading environment variables.")
		config.DBTYPE = os.Getenv("DB_TYPE")
		config.DatabaseURL = os.Getenv("DATABASE_URL")
		config.LinksColl = os.Getenv("LINKS_COLL")
		config.CounterColl = os.Getenv("COUNTER_COLL")
		config.Port = os.Getenv("PORT")
		config.RedirectMethod = os.Getenv("REDIRECT_METHOD")
		config.AuthToken = os.Getenv("AUTH_TOKEN")

		if len(config.DatabaseURL) == 0 {
			log.Fatalf("Missing config , Config : %+v", config)
		}

		return &config
	}

	configFile, err := os.Open(configFilename)
	defer configFile.Close()

	if configFile == nil || err != nil {
		log.Fatal("Failed to load config file", err)
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return &config
}
