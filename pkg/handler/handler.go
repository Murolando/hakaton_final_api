package handler

import (
	"github.com/Murolando/hakaton_final_api/pkg/service"
	"github.com/gin-gonic/gin"

)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
			auth.GET("/refresh/:refresh", h.newRefresh)
		}
		course := api.Group("/course")
		{
			course.GET("/",h.userIdentity,h.allCourse)
			course.GET("/:course-id",h.userIdentity,h.course)
		}
		finalTest := api.Group("/final")
		{
			finalTest.GET("/",h.userIdentity,h.startFinalTest)
		}
	}
	return router
}
