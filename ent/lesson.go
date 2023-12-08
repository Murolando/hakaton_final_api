package ent

type Lesson struct {
	LessonType string `json:"lesson-type"`
	CourseId   int    `json:"course-id"`
	CourseName string `json:"course-name"`
	Value      int    `json:"value"`
	Name       string `json:"name"`

	Questions []*Question `json:"questions,omitempty"`
	Materials []*Material `json:"materials,omitempty"`
}

type Question struct {
	QuestionType   string   `json:"question-type"`
	QuestionTypeId string   `json:"question-type-id"`
	LessonText     string   `json:"lesson-text"`
	Url            []string `json:"src-urls"`
	RightQuestion  int64
}

type Material struct {
	Name       string   `json:"name"`
	LessonText string   `json:"lesson-text"`
	Url        []string `json:"src-urls"`
}
