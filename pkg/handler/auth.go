package handler

import (
	"fmt"
	"net/http"

	"github.com/Murolando/hakaton_final_api/ent"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input ent.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(input)
	token, err := h.service.SignUp(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newResponse(c, "", token)
}

func (h *Handler) signIn(c *gin.Context) {
	var input ent.Auth

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.SignIn(input.Login, input.PasswordHash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	token, err := h.service.GenerateToken(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	refresh, err := h.service.NewRefreshToken(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "", map[string]interface{}{"token": token, "refresh": refresh})

}

func (h *Handler) newRefresh(c *gin.Context) {
	refresh := c.Param("refresh")

	id, err := h.service.GetByRefreshToken(refresh)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.service.GenerateToken(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newRefresh, err := h.service.NewRefreshToken(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, "", map[string]interface{}{"token": token, "refresh": newRefresh})
}
