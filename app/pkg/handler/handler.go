package handler

import (
	"secrets_keeper/app/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)

	router.LoadHTMLFiles("templates/index.html")
	router.GET("/", h.GetIndexPage)
	router.GET("/:key", h.GetMessage)

	return router
}
