package utils

import (
    "encoding/json"
    "net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
    RespondWithJSON(w, code, map[string]string{"message": message})
}

func RespondWithMessage(w http.ResponseWriter, message string) {
    RespondWithJSON(w, http.StatusOK, map[string]string{"message": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}
