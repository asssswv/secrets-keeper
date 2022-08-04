package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	//services *service.Service
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)

	router.LoadHTMLFiles("templates/index.html")
	router.GET("/", h.Home)

	return router
}
