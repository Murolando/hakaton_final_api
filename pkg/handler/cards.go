package handler

import (
	"net/http"
	"strconv"

	"github.com/Murolando/hakaton_final_api/ent"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addCard(c *gin.Context) {
	var input ent.Card
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userIdStr, _ := c.Get("userId")
	userId := userIdStr.(int64)
	rt, err := h.service.AddCard(&input, int(userId))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newResponse(c, "", rt)
}

func (h *Handler) deleteCard(c *gin.Context) {
	cardId, err := strconv.Atoi(c.Param("card-id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	isDeleted, err := h.service.DeleteCard(cardId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if !isDeleted {
		newErrorResponse(c, http.StatusBadRequest, "Information to delete was not found")
		return
	}
	newResponse(c, "", true)
}

func (h *Handler) getCard(c *gin.Context) {
	userIdStr, _ := c.Get("userId")
	userId := userIdStr.(int64)
	course, err := h.service.GetCard(int(userId))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newResponse(c, "", course)
}
