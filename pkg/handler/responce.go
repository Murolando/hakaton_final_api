package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Err struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	fmt.Println(message)
	c.AbortWithStatusJSON(statusCode, map[string]interface{}{
		"error":  map[string]interface{}{"code": statusCode, "msg": message},
		"result": false,
	})
}
func newResponse(c *gin.Context, str string, structure interface{}) {
	if str == "" {
		c.JSON(http.StatusOK, map[string]interface{}{
			"result": structure,
			"error":  map[string]int{"code": 200},
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"result": map[string]interface{}{str: structure},
			"error":  map[string]int{"code": 200},
		})
	}
}
