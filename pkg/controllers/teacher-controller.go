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

func GetTeachers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.Teachers)
}

func GetTeacher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, teacher := range config.Teachers {
		if teacher.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(teacher)
			return
		}
	}
	utils.RespondWithError(w, http.StatusNotFound, "Teacher not found")
}

func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher models.Teacher
	_ = json.NewDecoder(r.Body).Decode(&teacher)
	teacher.ID = strconv.Itoa(len(config.Teachers) + 1)
	config.Teachers = append(config.Teachers, teacher)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
}

func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, teacher := range config.Teachers {
		if teacher.ID == params["id"] {
			config.Teachers = append(config.Teachers[:index], config.Teachers[index+1:]...)
			var updatedTeacher models.Teacher
			_ = json.NewDecoder(r.Body).Decode(&updatedTeacher)
			updatedTeacher.ID = params["id"]
			config.Teachers = append(config.Teachers, updatedTeacher)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTeacher)
			return
		}
	}
	utils.RespondWithError(w, http.StatusNotFound, "Teacher not found")
}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, teacher := range config.Teachers {
		if teacher.ID == params["id"] {
			config.Teachers = append(config.Teachers[:index], config.Teachers[index+1:]...)
			utils.RespondWithMessage(w, "Teacher deleted successfully")
			return
		}
	}
	utils.RespondWithError(w, http.StatusNotFound, "Teacher not found")
}
