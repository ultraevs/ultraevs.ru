package router

import (
	"app/internal/api/controller"
)

func (router *Router) MainRoutes() {
	router.engine.GET("/", controller.GetMain)
	router.engine.GET("/gagarin", controller.GetSmolathon)
	router.engine.GET("/urbaton", controller.GetUrbaton)
	router.engine.GET("/moretech5", controller.GetMoretech5)
	router.engine.GET("/inno"), controoler.GetInnoHack)
}
