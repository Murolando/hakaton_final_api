package ent

type FinalTest struct {
	Id          int64          `json:"lesson-id"`
	MaxResult   int            `json:"max-result,omitempty"`
	LastResult   int            `json:"last-result,omitempty"`
	Description string         `json:"description,omitempty"`
	Question    []*FinalQuestion `json:"question"`
}

type FinalQuestion struct {
	Id             int64         `json:"material-id"`
	QuestionType   string        `json:"question-type"`
	QuestionTypeId string        `json:"question-type-id"`
	QuestText      string        `json:"quest-text"`
	Url            string        `json:"src-url,omitempty"`
	Answers        []*FinalAnswer `json:"answer"`
}

type FinalAnswer struct {
	Id         int64   `json:"answer-id"`
	Right      bool    `json:"right"`
	AnswerText string  `json:"answer-text"`
	QuestionId int64   `json:"question-id"`
	Url        *string `json:"src-url,omitempty"`
}

type UserAnswers struct{
	
}