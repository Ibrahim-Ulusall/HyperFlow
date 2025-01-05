package main

import (
	"HyperFlow/configurations"
	"fmt"
)

func main() {

	config := configurations.Config{}

	settings, err := config.GetConfiguration()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(settings.ConnectionString.DefaultConnection)
}
