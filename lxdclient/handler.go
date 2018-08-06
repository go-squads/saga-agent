package lxdclient

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lxc/lxd/shared/api"
)

// Handler ...
type Handler struct {
	client LxdClient
}

type createContainerRequestData struct {
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Server   string `json:"server,omitempty"`
	Alias    string `json:"alias,omitempty"`
}

type RequestUpdateStateContainer struct {
	Name  string                `json:"name"`
	State api.ContainerStatePut `json:"state"`
}

type deleteContainerRequestData struct {
	Name string `json:"name,omitempty"`
}

// GetContainersHandler ...
func (h *Handler) GetContainersHandler(w http.ResponseWriter, r *http.Request) {
	containers, err := h.client.getContainers()
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
	container, err := h.client.getContainer(vars["name"])
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

	op, err := h.client.createContainer(request)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, op)
	return
}

// DeleteContainerHandler ...
func (h *Handler) DeleteContainerHandler(w http.ResponseWriter, r *http.Request) {
	var data deleteContainerRequestData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	op, err := h.client.deleteContainer(data.Name)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, op)
	return
}

// UpdateStateContainerHandler ...
func (h *Handler) UpdateStateContainerHandler(w http.ResponseWriter, r *http.Request) {
	var data RequestUpdateStateContainer
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	op, err := h.client.updateContainerState(data.Name, data.State)
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
