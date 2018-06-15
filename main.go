package main

import (
	"fmt"
)

func main() {
	config := ReadConfig()
	fmt.Println(config)
	router := NewRouter(config)
	router.Run(":" + config.Port)
}
