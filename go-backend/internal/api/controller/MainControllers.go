package controller

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Message int `json:"code"`
}

// GetMainStat godoc
// @Summary Get Main
// @Description Get Main
// @Produce json
// @Tags main
// @Success 200 {object} SuccessResponse
// @Router / [get]
func GetMain(context *gin.Context) {
	response := SuccessResponse{Message: 200}
	context.HTML(200, "index.html", gin.H{"response": response})
}

func GetSmolathon(context *gin.Context) {
	response := SuccessResponse{Message: 200}
	context.HTML(200, "smolathon.html", gin.H{"response": response})
}
func GetUrbaton(context *gin.Context) {
	response := SuccessResponse{Message: 200}
	context.HTML(200, "urbaton.html", gin.H{"response": response})
}
func GetMoretech5(context *gin.Context) {
	response := SuccessResponse{Message: 200}
	context.HTML(200, "moretech5.html", gin.H{"response": response})
}
