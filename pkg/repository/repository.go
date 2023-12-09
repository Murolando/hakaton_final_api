package repository

import (
	"github.com/Murolando/hakaton_final_api/ent"
	"github.com/Murolando/hakaton_final_api/pkg/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	SignUp(user ent.User) (int64, error)
	GetUserByLoginAndPassword(mail *string, password *string) (int64, error)
	SetSession(user int64, refresh string, expiredAt string) error
	GetByRefreshToken(refresh string) (int64, error)
}
type Course interface {
	AllCourses(userId int) ([]*ent.ShortCourse, error)
	OneCourse(courseId int, userId int) (*ent.Course, error)
}
type FinalTestPst interface {
	StartFinalTest(userId int) (*ent.FinalTest, error)
}
type Card interface {
	AddCard(card *ent.Card, userId int) (bool, error)
	GetCard(userId int) ([]*ent.Card, error)
	DeleteCard(cardId int) (bool, error)
}

type Repository struct {
	Auth
	Course
	FinalTestPst
	Card
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:         postgres.NewAuthPostgres(db),
		Course:       postgres.NewCoursePostgres(db),
		FinalTestPst: postgres.NewFinalTestPostgres(db),
		Card:         postgres.NewCardPostgres(db),
	}
}
