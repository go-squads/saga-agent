package lxdclient

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lxc/lxd/shared/api"
)

// Handler ...
type Handler struct{}

type createContainerRequestData struct {
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Server   string `json:"server,omitempty"`
	Alias    string `json:"alias,omitempty"`
}

// GetContainersHandler ...
func (h *Handler) GetContainersHandler(w http.ResponseWriter, r *http.Request) {
	containers, err := getContainers()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, containers)
	return
}

// GetContainerHandler ...
func (h *Handler) GetContainerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	container, err := getContainer(vars["name"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, container)
	return
}

// CreateContainerHandler ...
func (h *Handler) CreateContainerHandler(w http.ResponseWriter, r *http.Request) {
	var data createContainerRequestData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	request := api.ContainersPost{
		Name: data.Name,
		Source: api.ContainerSource{
			Type:     data.Type,
			Protocol: data.Protocol,
			Server:   data.Server,
			Alias:    data.Alias,
		},
	}

	op, err := createContainer(request)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, op)
	return
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
