package handler

import (
	"fmt"
	"net/http"

	"github.com/Murolando/hakaton_final_api/ent"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body ent.UserRequest true "account info"
// @Success 200 {object} string
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/auth/sign-up [post]
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

// @Summary SignIn
// @Tags auth
// @Description auth in account
// @ID auth in account
// @Accept  json
// @Produce  json
// @Param input body ent.Auth true "email and password"
// @Success 200 {object} string
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/auth/sign-in [post]
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

// @Summary NewRefreshJwtTokens
// @Tags auth
// @Description generate new pair of jwt and refresh
// @ID gen new tokens
// @Accept  json
// @Produce  json
// @Param   refresh path string true "old refresh token"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router /api/auth/refresh/{refresh} [get]
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
