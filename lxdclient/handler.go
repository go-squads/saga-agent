package lxdclient

import (
	"encoding/json"
	"net/http"
)

// Handler ...
type Handler struct{}

// GetContainersHadler ...
func (h *Handler) GetContainersHadler(w http.ResponseWriter, R *http.Request) {
	containers, err := getContainers()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, containers)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
