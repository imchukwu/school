package config

import "github.com/imchukwu/school/pkg/models"

var Courses = []models.Course{
    {ID: 1, Name: "Mathematics", Description: "Basic Math course", Teacher: "John Doe"},
    {ID: 2, Name: "English", Description: "Basic English course", Teacher: "Jane Smith"},
}

var Classrooms = []models.Classroom{
    {ID: 1, Name: "Primary 2", Subclasses: []string{"Primary 2A", "Primary 2B"}},
    {ID: 2, Name: "Primary 3", Subclasses: []string{"Primary 3A", "Primary 3B"}},
}

var Students = []models.Student{
	{ID: "1", FirstName: "John", LastName: "Doe"},
	{ID: "2", FirstName: "Jane", LastName: "Smith"},
}

var Teachers = []models.Teacher{
	{ID: "1", FirstName: "Alice", LastName: "Johnson", Subject: "Math"},
	{ID: "2", FirstName: "Bob", LastName: "Brown", Subject: "Science"},
}

var Enrollments = []models.Enrollment{
	{ID: "1", StudentID: "1", ClassroomID: "101", CourseID: "201"},
	{ID: "2", StudentID: "2", ClassroomID: "102", CourseID: "202"},
}