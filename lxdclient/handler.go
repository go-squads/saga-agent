package lxdclient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler ...
type Handler struct{}

// GetContainersHandler ...
func (h *Handler) GetContainersHandler(w http.ResponseWriter, r *http.Request) {
	containers, err := getContainers()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, containers)
}

// GetContainerHandler ...
func (h *Handler) GetContainerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	container, err := getContainer(vars["name"])
	if err != nil {
		fmt.Println(err.Error())
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, container)
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
