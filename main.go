package main

import (
	"HyperFlow/configurations"
	"HyperFlow/web/routes"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	router := routes.Routes{}

	config := configurations.Config{}
	settings, error := config.ConfigureApp()
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	port := strconv.Itoa(settings.LaunchSettings.Port)
	fmt.Printf("Project is Running : http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, router.Invoke())

}
