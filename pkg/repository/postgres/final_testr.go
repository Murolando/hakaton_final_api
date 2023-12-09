package postgres

import (
	"errors"
	"fmt"

	"github.com/Murolando/hakaton_final_api/ent"
	"github.com/jmoiron/sqlx"
)

type FinalTestPostgres struct {
	db *sqlx.DB
}

func NewFinalTestPostgres(db *sqlx.DB) *FinalTestPostgres {
	return &FinalTestPostgres{
		db: db,
	}
}

func (r *FinalTestPostgres) StartFinalTest(userId int) (*ent.FinalTest, error) {
	var final ent.FinalTest

	// Id           int64  `json:"lesson-id"`
	// MaxResult    int    `json:"max-result,omitempty"`
	// MinResult    int    `json:"min-result,omitempty"`
	// Question *FinalQuestion `json:"question"`

	fmt.Println(1)
	query := fmt.Sprintf(`
	SELECT DISTINCT test_description
	FROM "%s"
	LIMIT 1`, finalTestTable)
	row := r.db.QueryRow(query)
	if err := row.Scan(&final.Description); err != nil {
		return nil, errors.New("bad Lesson Count")
	}

	fmt.Println(2)
	query = fmt.Sprintf(`
	SELECT DISTINCT max_result,last_result
	FROM "%s"
	WHERE user_id = $1
	LIMIT 1`, userFinalTable)
	row = r.db.QueryRow(query,userId)
	if err := row.Scan(&final.MaxResult,&final.LastResult); err != nil {
		if err.Error() != "sql: no rows in result set"{
			fmt.Println(err)
			return nil, errors.New("bad Lesson Count")
		}
		
	}
	fmt.Println(3)

	// lessons
	questions := make([]*ent.FinalQuestion, 0)
	query = fmt.Sprintf(`
	SELECT DISTINCT id,url,question,final_test_question_direction_id
	FROM "%s"`, finalTestQuestionTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next(){
		var quest ent.FinalQuestion
		rows.Scan(&quest.Id,&quest.Url,&quest.QuestText,&quest.QuestionTypeId)

		fmt.Println(4)
		q := fmt.Sprintf(`
		SELECT name
		FROM "%s" 
		WHERE id=$1`, finalTestQuestionDirectionTable)
		row2 := r.db.QueryRow(q, &quest.QuestionTypeId)
		if err := row2.Scan(&quest.QuestionType); err != nil {
			if err.Error() != "sql: no rows in result set"{
				fmt.Println(err)
				return nil, errors.New("bad question type")
			}
		}

		fmt.Println(5)
		
		// answers

		query = fmt.Sprintf(`
		SELECT DISTINCT answer_text,url,correct,final_test_question_id
		FROM "%s"
		WHERE final_test_question_id = $1`, finalTestAnswerTable)
		rows3, err := r.db.Query(query,quest.Id)
		if err != nil {
			return nil, err
		}
		for rows3.Next(){
			var ans ent.FinalAnswer
			rows3.Scan(&ans.AnswerText,&ans.Url,&ans.Right,&ans.QuestionId)
			quest.Answers = append(quest.Answers, &ans)
		}
		questions = append(questions, &quest)
	}
	final.Question = questions
	return &final, nil
}
