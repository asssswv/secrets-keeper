package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (h *Handler) GetMessage(c *gin.Context) {
	key := c.Param("key")

	message, err := h.services.Keeper.Get(key)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": message,
	})
}

func (h *Handler) SetMessage(c *gin.Context) {
	message := c.PostForm("message")
	key := h.services.KeyBuilder.Get()

	h.services.Keeper.Set(key, message)
	c.HTML(http.StatusCreated, "key.html", gin.H{"key": fmt.Sprintf("http://%s/%s", c.Request.Host, key)})
}