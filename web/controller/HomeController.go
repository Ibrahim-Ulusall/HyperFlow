package controller

import (
	core "HyperFlow/core/models"
	"HyperFlow/core/utilities/connection"
	"HyperFlow/core/utilities/helper"
	"HyperFlow/persistance/repositories"
	"fmt"
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

	data := core.PageDataModel{
		Title: "Anasayfa",
	}
	invoke()
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

func invoke() {
	db := connection.Database{}

	cnn, err := db.GetDbConnection()

	if err != nil {
		fmt.Printf("Hata : %v", err.Error())
	} else {

		userRepostiory := repositories.UserRepository{}
		userRepostiory.Database = cnn
		users, _ := userRepostiory.GetList()

		for _, item := range users {
			fmt.Printf("id : %s username: %s\n", item.ID, item.Username)
		}

	}

}
