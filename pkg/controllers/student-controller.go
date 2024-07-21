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

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.Students)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, student := range config.Students {
		if student.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	utils.RespondWithError(w, http.StatusNotFound, "Student not found")
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	student.ID = strconv.Itoa(len(config.Students) + 1)
	config.Students = append(config.Students, student)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, student := range config.Students {
		if student.ID == params["id"] {
			config.Students = append(config.Students[:index], config.Students[index+1:]...)
			var updatedStudent models.Student
			_ = json.NewDecoder(r.Body).Decode(&updatedStudent)
			updatedStudent.ID = params["id"]
			config.Students = append(config.Students, updatedStudent)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedStudent)
			return
		}
	}
	utils.RespondWithError(w, http.StatusNotFound, "Student not found")
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, student := range config.Students {
		if student.ID == params["id"] {
			config.Students = append(config.Students[:index], config.Students[index+1:]...)
			utils.RespondWithMessage(w, "Student deleted successfully")
			return
		}
	}
	utils.RespondWithError(w, http.StatusNotFound, "Student not found")
}
