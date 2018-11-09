package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	baseconv, _    = NewBaseConvertor(62)
	errInvalidLink = errors.New("short link too large")
	config         *Configuration
)

func main() {
	config = ReadConfig()
	fmt.Println(config)

	var (
		counterRepo CounterRepository
		linkRepo    LinkRepository
	)

	if config.DBTYPE == "mongodb" {
		mongodb := NewMongoDBRepository()
		linkRepo = mongodb
		counterRepo = mongodb
	} else if config.DBTYPE == "postgresql" {
		postgresql := NewPostgreSQLRepository()
		linkRepo = postgresql
		counterRepo = postgresql
	} else {
		log.Fatal("Unrecognized DBTYPE")
	}

	service := &ServiceImpl{linkRepo, counterRepo}
	router := NewRouter(service)
	router.Run(":" + config.Port)
	defer service.close()
}
