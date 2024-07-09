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

func GetClassrooms(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(config.Classrooms)
}

func GetClassroom(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for _, classroom := range config.Classrooms {
        if classroom.ID == id {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(classroom)
            return
        }
    }
    utils.RespondWithError(w, http.StatusNotFound, "Classroom not found")
}

func CreateClassroom(w http.ResponseWriter, r *http.Request) {
    var classroom models.Classroom
    _ = json.NewDecoder(r.Body).Decode(&classroom)
    classroom.ID = len(config.Classrooms) + 1
    config.Classrooms = append(config.Classrooms, classroom)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(classroom)
}

func UpdateClassroom(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for index, classroom := range config.Classrooms {
        if classroom.ID == id {
            config.Classrooms = append(config.Classrooms[:index], config.Classrooms[index+1:]...)
            var updatedClassroom models.Classroom
            _ = json.NewDecoder(r.Body).Decode(&updatedClassroom)
            updatedClassroom.ID = id
            config.Classrooms = append(config.Classrooms, updatedClassroom)
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(updatedClassroom)
            return
        }
    }
    utils.RespondWithError(w, http.StatusNotFound, "Classroom not found")
}

func DeleteClassroom(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for index, classroom := range config.Classrooms {
        if classroom.ID == id {
            config.Classrooms = append(config.Classrooms[:index], config.Classrooms[index+1:]...)
            utils.RespondWithMessage(w, "Classroom deleted successfully")
            return
        }
    }
    utils.RespondWithError(w, http.StatusNotFound, "Classroom not found")
}
