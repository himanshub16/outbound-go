package main

// Credits for this file: https://www.thepolyglotdeveloper.com/2017/04/load-json-configuration-file-golang-application/

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration describes the fields of JSON configuration
type Configuration struct {
	DBURL       string `json:"DB_URL"`
	DBName      string `json:"DB_NAME"`
	LinksColl   string `json:"LINKS_COLL"`
	CounterColl string `json:"COUNTER_COLL"`

	Port           string `json:"PORT"`
	RedirectMethod string `json:"REDIRECT_METHOD"`
	AuthToken      string `json:"AUTH_TOKEN"`
}

// ReadConfig reads config from config file
func ReadConfig() *Configuration {
	var config Configuration
	var configFilename string

	configFilename = os.Getenv("CONFIG_FILE")
	if len(configFilename) == 0 {
		// configFilename = ".env.json"
		log.Println("Config file not provided. Reading environment variables.")
		config.DBName = os.Getenv("DB_NAME")
		config.DBURL = os.Getenv("DB_URL")
		config.LinksColl = os.Getenv("LINKS_COLL")
		config.CounterColl = os.Getenv("COUNTER_COLL")
		config.Port = os.Getenv("PORT")
		config.RedirectMethod = os.Getenv("REDIRECT_METHOD")
		config.AuthToken = os.Getenv("AUTH_TOKEN")

		if len(config.DBURL) == 0 || len(config.DBName) == 0 {
			log.Fatal("Missing config")
		}

		return &config
	}

	configFile, err := os.Open(configFilename)
	defer configFile.Close()

	if err != nil {
		log.Fatal("Failed to load config file", err)
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return &config
}
