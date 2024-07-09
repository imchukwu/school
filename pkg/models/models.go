package models

type Course struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Teacher     string `json:"teacher"`
}

type Classroom struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Subclasses []string `json:"subclasses"`
}
