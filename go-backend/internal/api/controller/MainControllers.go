package controller

import (
	_ "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type CodeResponse struct {
	Message int `json:"code"`
}

type CookieResponse struct {
	Cookie string `json:"cookie"`
}

// GetMain godoc
// @Summary Get Main
// @Description Get Main
// @Produce json
// @Tags main
// @Success 200 {object} CookieResponse
// @Router / [get]
func GetMain(context *gin.Context) {
	cookie, err := context.Cookie("session")
	if err != nil {
		response := CookieResponse{Cookie: "Zero"}
		context.HTML(401, "index.html", gin.H{"response": response})
		return
	}
	response := CookieResponse{Cookie: cookie}
	context.HTML(200, "index.html", gin.H{"response": response})
}

// GetSmolathon godoc
// @Summary Get Smolathon
// @Description Get Smolathon
// @Produce json
// @Tags Smolathon
// @Success 200 {object} CodeResponse
// @Router /smolathon [get]
func GetSmolathon(context *gin.Context) {
	response := CodeResponse{Message: 200}
	context.HTML(200, "smolathon.html", gin.H{"response": response})
}

// GetUrbaton godoc
// @Summary Get Urbaton
// @Description Get Urbaton
// @Produce json
// @Tags Urbaton
// @Success 200 {object} CodeResponse
// @Router /urbaton [get]
func GetUrbaton(context *gin.Context) {
	response := CodeResponse{Message: 200}
	context.HTML(200, "urbaton.html", gin.H{"response": response})
}

// GetMoretech5 godoc
// @Summary Get Moretech5
// @Description Get Moretech5
// @Produce json
// @Tags Moretech5
// @Success 200 {object} CodeResponse
// @Router /moretech5 [get]
func GetMoretech5(context *gin.Context) {
	response := CodeResponse{Message: 200}
	context.HTML(200, "moretech5.html", gin.H{"response": response})
}
