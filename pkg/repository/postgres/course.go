package postgres

import (
	"errors"
	"fmt"

	"github.com/Murolando/hakaton_final_api/ent"
	"github.com/jmoiron/sqlx"
)

type CoursePostgres struct {
	db *sqlx.DB
}

func NewCoursePostgres(db *sqlx.DB) *CoursePostgres {
	return &CoursePostgres{
		db: db,
	}
}

func (r *CoursePostgres) AllCourses(userId int) ([]*ent.ShortCourse, error) {
	fmt.Println(1)
	courses := make([]*ent.ShortCourse, 0)
	query := fmt.Sprintf(`
	SELECT DISTINCT id,description,name,course_age,url
	FROM "%s"`, courseTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var sh ent.ShortCourse
		if err := rows.Scan(&sh.Id, &sh.Description, &sh.Name, &sh.CourseDifficulty, &sh.Url); err != nil {
			return nil, err
		}
		fmt.Println(2)
		var allLessonCount int
		var userLessonCount int
		q := fmt.Sprintf(`
		SELECT COUNT(id)
		FROM "%s" 
		WHERE course_id=$1 `, lessonTable)
		row := r.db.QueryRow(q, sh.Id)
		if err := row.Scan(&allLessonCount); err != nil {
			return nil, errors.New("bad Lesson Count")
		}
		fmt.Println(3)
		q = fmt.Sprintf(`
		SELECT COUNT(id)
		FROM "%s" 
		WHERE user_id=$1 AND course_id = $2 `, userLessonTable)
		row = r.db.QueryRow(q, userId, sh.Id)
		if err := row.Scan(&userLessonCount); err != nil {
			return nil, errors.New("bad UserLesson Count")
		}

		p := int(float32(userLessonCount) / float32(allLessonCount) * 100)
		sh.Progress = p
		courses = append(courses, &sh)

	}
	fmt.Println(4)
	return courses, nil
}
func (r *CoursePostgres) OneCourse(courseId int, userId int) (*ent.Course, error) {
	fmt.Println(1,courseId,userId)
	var course ent.Course
	query := fmt.Sprintf(`
	SELECT DISTINCT description,name,course_age,url
	FROM "%s" 
	WHERE id=$1`, courseTable)
	row := r.db.QueryRow(query, courseId)
	if err := row.Scan(&course.Description, &course.Name, &course.CourseDifficulty, &course.Url); err != nil {
		return nil, errors.New("bad Lesson Count")
	}
	fmt.Println(2)
	// progress
	var allLessonCount int
	var userLessonCount int
	q := fmt.Sprintf(`
	SELECT COUNT(id)
	FROM "%s" 
	WHERE course_id=$1 `, lessonTable)
	row = r.db.QueryRow(q, courseId)
	if err := row.Scan(&allLessonCount); err != nil {
		return nil, errors.New("bad Lesson Count")
	}
	fmt.Println(3)
	q = fmt.Sprintf(`
	SELECT COUNT(id)
	FROM "%s" 
	WHERE user_id=$1 AND course_id = $2 `, userLessonTable)
	row = r.db.QueryRow(q, userId, courseId)
	if err := row.Scan(&userLessonCount); err != nil {
		return nil, errors.New("bad UserLesson Count")
	}

	p := int(float32(userLessonCount) / float32(allLessonCount) * 100)
	course.Progress = p
	fmt.Println(4)
	// lessons
	lessons := make([]*ent.Lesson, 0)
	query = fmt.Sprintf(`
	SELECT DISTINCT id,lesson_type,course_id,value,name
	FROM "%s"
	WHERE course_id = $1`, lessonTable)
	rows, err := r.db.Query(query, courseId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var l ent.Lesson
		if err := rows.Scan(&l.Id, &l.LessonTypeId, &l.CourseId, &l.Value, &l.Name); err != nil {
			return nil, err
		}
		if l.LessonTypeId == 1 {
			fmt.Println(5)
			//material
			var m ent.Material
			q = fmt.Sprintf(`
			SELECT id,name,lesson_text
			FROM "%s" 
			WHERE lesson_id=$1`, lessonMatTable)
			row = r.db.QueryRow(q, l.Id)
			if err := row.Scan(&m.Id, &m.Name, &m.LessonText); err != nil {
				fmt.Println(err)
				return nil, errors.New("bad material")
			}
			fmt.Println(6)
			// material urls
			q = fmt.Sprintf(`
			SELECT url
			FROM "%s" 
			WHERE lesson_mat_id=$1`, lessonMatSrcTable)
			rows2, err := r.db.Query(q, m.Id)
			if err != nil {
				return nil, err
			}
			for rows2.Next() {
				var url string
				rows2.Scan(&url)
				m.Url = append(m.Url, url)
			}
			if err := row.Scan(&m.Id, &m.Name, &m.LessonText); err != nil {
				if err.Error() != "sql: no rows in result set"{
					fmt.Println(err)
					return nil, errors.New("bad material urls")
				}
			}
			l.Material = &m
		} else {
			fmt.Println(7)
			//question
			var question ent.Question
			q = fmt.Sprintf(`
			SELECT id,question,lesson_test_question_type_id,url
			FROM "%s" 
			WHERE lesson_id=$1`, lessonTestQuestionTable)
			row = r.db.QueryRow(q, l.Id)
			if err := row.Scan(&question.Id, &question.QuestText, &question.QuestionTypeId,&question.Url); err != nil {
				fmt.Println(err)
				return nil, errors.New("bad question")
			}
			fmt.Println(8)
			// type name
			q = fmt.Sprintf(`
			SELECT name
			FROM "%s" 
			WHERE id=$1`, lessonTestQuestionTypeTable)
			row = r.db.QueryRow(q, question.QuestionTypeId)
			if err := row.Scan(&question.QuestionType); err != nil {
				fmt.Println(err)
				return nil, errors.New("bad question type")
			}
			fmt.Println(9)
			// answers
			qur := fmt.Sprintf(`
			SELECT id,answer_text,correct
			FROM "%s" 
			WHERE lesson_test_question_id=$1`, lessonTestAnswerTable)
			rows2, err := r.db.Query(qur, question.Id)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			for rows2.Next() {
				var ans ent.Answer
				var url string
				ans.QuestionId = question.Id
				rows2.Scan(&ans.Id,&ans.AnswerText, &ans.Right)
				fmt.Println(10)
				q = fmt.Sprintf(`
				SELECT url
				FROM "%s" 
				WHERE lesson_test_answer_id=$1`, lessonTestAnswerSrcTable)
				row3 := r.db.QueryRow(q, ans.Id)
				if err := row3.Scan(&url); err != nil {
					if err.Error() != "sql: no rows in result set"{
						fmt.Println(err)
						return nil, errors.New("bad question type")
					}
					
				}
				ans.Url = &url
				question.Answers = append(question.Answers, ans)
			}
			l.Question = &question

		}
		lessons = append(lessons, &l)

	}

	course.Lessons = lessons
	return &course, nil
}
