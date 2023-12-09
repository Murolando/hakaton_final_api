package service

import (
	"github.com/Murolando/hakaton_final_api/ent"
	"github.com/Murolando/hakaton_final_api/pkg/repository"
)

type Auth interface {
	SignIn(mail *string, password *string) (int64, error)
	SignUp(user ent.User) (map[string]interface{}, error)
	GenerateToken(id int64) (string, error)
	ParseToken(accesstoken string) (int64, error)
	NewRefreshToken(id int64) (string, error)
	GetByRefreshToken(refresh string) (int64, error)
}
type Course interface {
	AllCourses(userId int)([]*ent.ShortCourse,error)
	OneCourse(courseId int,userId int )(*ent.Course,error)
}
type Service struct {
	Auth
	Course
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
		Course: NewCourseService(repo),
	}
}
