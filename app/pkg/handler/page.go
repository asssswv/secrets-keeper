package handler

import (
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
