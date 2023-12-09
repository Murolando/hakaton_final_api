package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_"github.com/lib/pq"
)

type Config struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

const (
	userTable = "user"
	courseTable ="course"
	forwardTable = "forward"
	forwardSrcTable = "forward_src"
	lessonTable = "lesson"
	lessonMatTable = "lesson_mat"
	lessonMatSrcTable = "lesson_mat_src"
	lessonTestTable = "lesson_test"
	lessonTestAnswerTable = "lesson_test_answer"
	lessonTestAnswerSrcTable = "lesson_test_answer_src"
	lessonTestQuestionTable = "lesson_test_question"
	lessonTestQuestionTypeTable = "lesson_test_question_type"
	lessonTypeTable = "lesson_type"
	roleTable = "role"
	userLessonTable = "user_lesson"

)

func NewConfig(host string, port string, user string, password string, dbname string) *Config {
	return &Config{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
	}
}

func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host = %s port = %s user = %s dbname = %s password = %s  sslmode = disable",
		cfg.host, cfg.port, cfg.user, cfg.dbname, cfg.password)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
