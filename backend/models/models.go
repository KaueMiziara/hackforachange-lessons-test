package models

type Grade struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Subjects []Subject `json:"subjects"`
}

type Subject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
