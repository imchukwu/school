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

type Student struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Teacher struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Subject   string `json:"subject"`
}

type Enrollment struct {
	ID         string `json:"id"`
	StudentID  string `json:"student_id"`
	ClassroomID string `json:"classroom_id"`
	CourseID   string `json:"course_id"`
}