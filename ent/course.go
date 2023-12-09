package ent

type Course struct {
	Description      string `json:"description"`
	CourseDifficulty int8   `json:"course-difficulty"`
	Url              string `json:"url"`
	Name             string `json:"name"`

	Progress int       `json:"progress"`
	Lessons  []*Lesson `json:"lessons"`
}

type ShortCourse struct {
	Id               int64  `json:"course-id"`
	Description      string `json:"description"`
	CourseDifficulty int8   `json:"course-difficulty"`
	Url              string `json:"url"`
	Name             string `json:"name"`
	Progress         int    `json:"progress"`
}
