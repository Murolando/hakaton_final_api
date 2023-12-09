package service

import (
	"github.com/Murolando/hakaton_final_api/ent"
	"github.com/Murolando/hakaton_final_api/pkg/repository"
)

type CourseService struct {
	repo *repository.Repository
}

func NewCourseService(repo *repository.Repository) *CourseService {
	return &CourseService{
		repo: repo,
	}
}

func (s *CourseService) AllCourses(userId int) ([]*ent.ShortCourse, error) {

	courses,err := s.repo.AllCourses(userId)
	if err!=nil{
		return nil,err
	}
	return courses,nil
}
func (s *CourseService) OneCourse(courseId int, userId int) (*ent.Course, error) {
	course,err := s.repo.OneCourse(courseId,userId)
	if err!=nil{
		return nil,err
	}
	return course,nil
}
