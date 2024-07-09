package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/imchukwu/school/pkg/config"
    "github.com/imchukwu/school/pkg/models"
    "github.com/imchukwu/school/pkg/utils"
)

func GetCourses(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(config.Courses)
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for _, course := range config.Courses {
        if course.ID == id {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(course)
            return
        }
    }
    utils.RespondWithError(w, http.StatusNotFound, "Course not found")
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
    var course models.Course
    _ = json.NewDecoder(r.Body).Decode(&course)
    course.ID = len(config.Courses) + 1
    config.Courses = append(config.Courses, course)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(course)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for index, course := range config.Courses {
        if course.ID == id {
            config.Courses = append(config.Courses[:index], config.Courses[index+1:]...)
            var updatedCourse models.Course
            _ = json.NewDecoder(r.Body).Decode(&updatedCourse)
            updatedCourse.ID = id
            config.Courses = append(config.Courses, updatedCourse)
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(updatedCourse)
            return
        }
    }
    utils.RespondWithError(w, http.StatusNotFound, "Course not found")
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for index, course := range config.Courses {
        if course.ID == id {
            config.Courses = append(config.Courses[:index], config.Courses[index+1:]...)
            utils.RespondWithMessage(w, "Course deleted successfully")
            return
        }
    }
    utils.RespondWithError(w, http.StatusNotFound, "Course not found")
}
