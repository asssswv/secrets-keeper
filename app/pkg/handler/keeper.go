package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"maxTTL": 86400, "maxMessageLength": 1024})
}

func (h *Handler) GetMessage(c *gin.Context) {
	key := c.Param("key")

	message, err := h.services.Keeper.Get(key)

	if err != nil {

		if err.Error() == "message not found" {
			c.HTML(http.StatusNotFound, "404.html", gin.H{})
			newErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		c.HTML(http.StatusInternalServerError, "500.html", gin.H{})
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "message.html", gin.H{"message": message})
}

func (h *Handler) SetMessage(c *gin.Context) {
	message := c.PostForm("message")
	ttl := c.PostForm("ttl")

	key, err := h.services.UUIDKeyBuilder.Get()

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{})
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	t, err := strconv.Atoi(ttl)
	if err != nil {
		c.HTML(http.StatusBadRequest, "400.html", gin.H{})
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Keeper.Set(key, message, t)
	if err != nil {
		if err.Error() == "message length too long!" || err.Error() == "ttl too long" {
			c.HTML(http.StatusBadRequest, "400.html", gin.H{})
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.HTML(http.StatusInternalServerError, "500.html", gin.H{})
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusCreated, "key.html", gin.H{"key": fmt.Sprintf("http://%s/message/%s", c.Request.Host, key)})
}
