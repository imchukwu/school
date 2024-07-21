package routes

import (
	"github.com/gorilla/mux"
	"github.com/imchukwu/school/pkg/controllers"
)

func SetupRouter() *mux.Router {
    router := mux.NewRouter()

    // Course Management Routes
    router.HandleFunc("/api/courses", controllers.GetCourses).Methods("GET")
    router.HandleFunc("/api/courses/{id}", controllers.GetCourse).Methods("GET")
    router.HandleFunc("/api/courses", controllers.CreateCourse).Methods("POST")
    router.HandleFunc("/api/courses/{id}", controllers.UpdateCourse).Methods("PUT")
    router.HandleFunc("/api/courses/{id}", controllers.DeleteCourse).Methods("DELETE")

    // Classroom Management Routes
    router.HandleFunc("/api/classrooms", controllers.GetClassrooms).Methods("GET")
    router.HandleFunc("/api/classrooms/{id}", controllers.GetClassroom).Methods("GET")
    router.HandleFunc("/api/classrooms", controllers.CreateClassroom).Methods("POST")
    router.HandleFunc("/api/classrooms/{id}", controllers.UpdateClassroom).Methods("PUT")
    router.HandleFunc("/api/classrooms/{id}", controllers.DeleteClassroom).Methods("DELETE")

     // Student Management Routes
    router.HandleFunc("/api/students", controllers.GetStudents).Methods("GET")
	router.HandleFunc("/api/students/{id}", controllers.GetStudent).Methods("GET")
	router.HandleFunc("/api/students", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/api/students/{id}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/api/students/{id}", controllers.DeleteStudent).Methods("DELETE")

	// Teacher Management routes
	router.HandleFunc("/api/teachers", controllers.GetTeachers).Methods("GET")
	router.HandleFunc("/api/teachers/{id}", controllers.GetTeacher).Methods("GET")
	router.HandleFunc("/api/teachers", controllers.CreateTeacher).Methods("POST")
	router.HandleFunc("/api/teachers/{id}", controllers.UpdateTeacher).Methods("PUT")
	router.HandleFunc("/api/teachers/{id}", controllers.DeleteTeacher).Methods("DELETE")

	// Enrollment Management routes
	router.HandleFunc("/api/enrollments", controllers.GetEnrollments).Methods("GET")
	router.HandleFunc("/api/enrollments/{id}", controllers.GetEnrollment).Methods("GET")
	router.HandleFunc("/api/enrollments", controllers.CreateEnrollment).Methods("POST")
	router.HandleFunc("/api/enrollments/{id}", controllers.UpdateEnrollment).Methods("PUT")
	router.HandleFunc("/api/enrollments/{id}", controllers.DeleteEnrollment).Methods("DELETE")

    return router
}
