package ent

type Lesson struct {
	Id           int64  `json:"lesson-id"`
	LessonTypeId int    `json:"lesson-type-id"`
	CourseId     int    `json:"course-id"`
	CourseName   string `json:"course-name"`
	Value        int    `json:"value"`
	Name         string `json:"name"`
	Passed       bool   `json:"passed"`

	Question *Question `json:"question,omitempty"`
	Material *Material `json:"material,omitempty"`
}

type Question struct {
	Id             int64    `json:"material-id"`
	QuestionType   string   `json:"question-type"`
	QuestionTypeId string   `json:"question-type-id"`
	QuestText      string   `json:"quest-text"`
	Url            string   `json:"src-url,omitempty"`
	Answers        []Answer `json:"answer"`
}

type Material struct {
	Id         int64    `json:"material-id"`
	Name       *string   `json:"name"`
	LessonText string   `json:"lesson-text"`
	Url        []string `json:"src-urls,omitempty"`
}

type Answer struct {
	Id         int64   `json:"answer-id"`
	Right      bool    `json:"right"`
	AnswerText string  `json:"answer-text"`
	QuestionId int64   `json:"question-id"`
	Url        *string `json:"src-url,omitempty"`
}
