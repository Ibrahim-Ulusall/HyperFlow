package controller

import (
	core "HyperFlow/core/models"
	"HyperFlow/core/utilities/helper"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HomeController struct {
}

func (homeController HomeController) HandlerIndex(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	view, error := template.ParseFiles(helper.Include("home")...)
	if error != nil {
		log.Println(error.Error())
		http.Error(responseWriter, "File Parse Error", http.StatusInternalServerError)
		return
	}

	data := core.PageDataModel{
		Title: "Anasayfa",
	}
	view.ExecuteTemplate(responseWriter, "layout", data)
}

func (homeController HomeController) HandlerPageNotFound(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	view, error := template.ParseFiles(helper.Include("home")...)
	if error != nil {
		log.Println(error.Error())
		http.Error(responseWriter, "File Parse Error", http.StatusInternalServerError)
		return
	}
	view.ExecuteTemplate(responseWriter, "pageNotFound", nil)
}
