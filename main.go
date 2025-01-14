package main

import (
	"HyperFlow/web/routes"
	"net/http"
)

func main() {
	http.ListenAndServe(":80", routes.Routes{}.Invoke())
}
