package routes

import (
	"HyperFlow/web/controller"

	"github.com/julienschmidt/httprouter"
)

type Routes struct {
}

func (route Routes) Invoke() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", controller.HomeController{}.HandlerIndex)

	return router
}
