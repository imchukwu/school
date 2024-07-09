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
