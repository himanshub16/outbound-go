package main

import (
	"fmt"
	"log"
)

func main() {
	config := ReadConfig()
	fmt.Println(config)
	router := NewRouter(config)
	if err := router.Run(":" + config.Port); err != nil {
		log.Fatalf("Failed running the service with Config : %+v ,Error: %+v ", config, err)
	}
}
