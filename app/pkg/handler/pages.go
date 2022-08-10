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

		if err.Error() == "message not found" {
			c.HTML(http.StatusNotFound, "404.html", gin.H{})
			return
		}

		c.HTML(http.StatusInternalServerError, "500.html", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "message.html", gin.H{"message": message})
}

func (h *Handler) SetMessage(c *gin.Context) {
	message := c.PostForm("message")
	key, err := h.services.UUIDKeyBuilder.Get()

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{})
		return
	}

	err = h.services.Keeper.Set(key, message)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{})
		return
	}

	c.HTML(http.StatusCreated, "key.html", gin.H{"key": fmt.Sprintf("http://%s/message/%s", c.Request.Host, key)})
}
