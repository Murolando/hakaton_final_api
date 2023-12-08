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
	course ="coruse"
	forward = "forward"
	forwardSrc = "forward_src"
	lesson = "lesson"
	lessonMat = "lesson_mat"
	lessonMatSrc = "lesson_mat_src"
	lessonTest = "lesson_test"
	lessonTestAnswer = "lesson_test_answer"
	lessonTestAnswerSrc = "lesson_test_answer_src"
	lessonTestQuestion = "lesson_test_question"
	lessonTestQuestionType = "lesson_test_question_type"
	lessonType = "lesson_type"
	role = "role"
	userLesson = "user_lesson"

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
