package routes

import (
	"HyperFlow/web/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Routes struct {
}

func (route Routes) Invoke() *httprouter.Router {
	router := httprouter.New()

	router.ServeFiles("/assets/*filepath", http.Dir("web/assets"))
	routeHome(router)
	return router
}

func routeHome(router *httprouter.Router) {
	home := controller.HomeController{}
	router.GET("/", home.Index)
	router.GET("/pageNotFound", home.PageNotFound)

}
