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

func GetEnrollments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.Enrollments)
}

func GetEnrollment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, enrollment := range config.Enrollments {
		if enrollment.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(enrollment)
			return
		}
	}
	utils.RespondWithError(w, http.StatusNotFound, "Enrollment not found")
}

func CreateEnrollment(w http.ResponseWriter, r *http.Request) {
	var enrollment models.Enrollment
	_ = json.NewDecoder(r.Body).Decode(&enrollment)
	enrollment.ID = strconv.Itoa(len(config.Enrollments) + 1)
	config.Enrollments = append(config.Enrollments, enrollment)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enrollment)
}

func UpdateEnrollment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, enrollment := range config.Enrollments {
		if enrollment.ID == params["id"] {
			config.Enrollments = append(config.Enrollments[:index], config.Enrollments[index+1:]...)
			var updatedEnrollment models.Enrollment
			_ = json.NewDecoder(r.Body).Decode(&updatedEnrollment)
			updatedEnrollment.ID = params["id"]
			config.Enrollments = append(config.Enrollments, updatedEnrollment)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedEnrollment)
			return
		}
	}
	utils.RespondWithError(w, http.StatusNotFound, "Enrollment not found")
}

func DeleteEnrollment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, enrollment := range config.Enrollments {
		if enrollment.ID == params["id"] {
			config.Enrollments = append(config.Enrollments[:index], config.Enrollments[index+1:]...)
			utils.RespondWithMessage(w, "Enrollment deleted successfully")
			return
		}
	}
	utils.RespondWithError(w, http.StatusNotFound, "Enrollment not found")
}
