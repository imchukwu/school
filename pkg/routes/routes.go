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

    return router
}
