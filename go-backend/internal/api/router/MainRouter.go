package router

import (
	"app/internal/api/controller"
)

func (router *Router) MainRoutes() {
	router.engine.GET("/", controller.GetMain)
	router.engine.GET("/smolathon", controller.GetSmolathon)
	router.engine.GET("/urbaton", controller.GetUrbaton)
	router.engine.GET("/moretech5", controller.GetMoretech5)
}
