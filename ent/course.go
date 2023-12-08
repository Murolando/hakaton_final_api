package ent

type Course struct {
	Description string `json:"description"`
	CourseAge   int8   `json:"course_age"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Progress    int    `json:"progress"`

	Lessons []*Lesson `json:"lessons"`
}

type ShortCourse struct {
	Description string `json:"description"`
	CourseAge   int8   `json:"course_age"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Progress    int    `json:"progress"`
}
