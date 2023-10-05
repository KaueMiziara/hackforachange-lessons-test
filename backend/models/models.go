package models

type JsonData struct {
	Grades []Grade `json:"grades"`
}

type Grade struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Subjects []Subject `json:"subjects"`
}

type Subject struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Lessons   []Lesson   `json:"lessons"`
	Exercises []Exercise `json:"exercises"`
}

type Lesson struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Exercise struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
