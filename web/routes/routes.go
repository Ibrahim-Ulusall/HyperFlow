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

	router.GET("/", controller.HomeController{}.HandlerIndex)
	router.GET("/pageNotFound", controller.HomeController{}.HandlerPageNotFound)

	return router
}
