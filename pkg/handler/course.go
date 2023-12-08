package handler

import (
	"fmt"
	"net/http"

	"github.com/Murolando/hakaton_final_api/ent"
	"github.com/gin-gonic/gin"
)

func (h *Handler) course(c *gin.Context) {
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

func (h *Handler) AllCourse(c *gin.Context) {
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
