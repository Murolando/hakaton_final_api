package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) startFinalTest(c *gin.Context) {
	userIdStr, _ := c.Get("userId")
	userId := userIdStr.(int64)
	finalTest, err := h.service.StartFinalTest(int(userId))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newResponse(c, "", finalTest)
}
