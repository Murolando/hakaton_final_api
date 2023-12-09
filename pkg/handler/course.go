package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) course(c *gin.Context) {
	courseId, err := strconv.Atoi(c.Param("course-id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userIdStr, _ := c.Get("userId")
	userId := userIdStr.(int64)
	course, err := h.service.OneCourse(courseId, int(userId))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newResponse(c, "", course)
}

func (h *Handler) allCourse(c *gin.Context) {
	userIdStr, _ := c.Get("userId")

	userId := userIdStr.(int64)
	courses, err := h.service.AllCourses(int(userId))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newResponse(c, "courses", courses)
}

// func (h *Handler) completeLesson(c *gin.Context) {
// 	lessonId, err := strconv.Atoi(c.Param("lesson-id"))
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	userIdStr, _ := c.Get("userId")
// 	userId := userIdStr.(int64)
// 	courses, err := h.service.AllCourses(int(userId))
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	newResponse(c, "courses", courses)
// }
