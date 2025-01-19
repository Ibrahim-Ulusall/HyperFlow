package controller

import (
	core_models "HyperFlow/core/models"
	"HyperFlow/core/utilities/helper"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HomeController struct {
}

func (homeController HomeController) Index(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	view, error := template.ParseFiles(helper.Include("home")...)
	if error != nil {
		log.Println(error.Error())
		http.Error(responseWriter, "File Parse Error", http.StatusInternalServerError)
		return
	}

	data := core_models.PageDataModel{
		Title: "Anasayfa",
	}
	view.ExecuteTemplate(responseWriter, "layout", data)
}

func (homeController HomeController) PageNotFound(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	view, error := template.ParseFiles(helper.Include("home")...)
	if error != nil {
		log.Println(error.Error())
		http.Error(responseWriter, "File Parse Error", http.StatusInternalServerError)
		return
	}
	view.ExecuteTemplate(responseWriter, "pageNotFound", nil)
}
