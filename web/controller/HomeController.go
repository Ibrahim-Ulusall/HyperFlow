package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HomeController struct {
}

func (homeController HomeController) HandlerIndex(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	body := request.Body
	fmt.Println(body)
	parameters := params

	for p := range parameters {
		fmt.Println(p)
	}

	responseWriter.Write([]byte("Hi"))
}
