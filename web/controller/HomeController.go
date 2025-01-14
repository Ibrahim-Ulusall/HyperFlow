package controller

import (
	"HyperFlow/core/utilities/helper"
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HomeController struct {
}

func (homeController HomeController) HandlerIndex(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	view, error := template.ParseFiles(helper.Include("home")...)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	view.ExecuteTemplate(responseWriter, "home", nil)
}

func (homeController HomeController) HandlerPageNotFound(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	view, error := template.ParseFiles(helper.Include("home")...)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	view.ExecuteTemplate(responseWriter, "pageNotFound", nil)
}
