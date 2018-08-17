package main

import (
	"log"
	"net/http"

	"github.com/go-squads/saga-agent/lxdclient"
	"github.com/gorilla/mux"
)

// App ...
type App struct {
	Router *mux.Router
}

// Run ...
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

// Initialize ...
func (a *App) Initialize() {
	a.Router = mux.NewRouter()

	handler := lxdclient.Handler{}
	handler.HandlerClient = &lxdclient.LxdClient{}
	handler.HandlerClient.Init()

	worker := CronWorker{}
	worker.CronClient = &lxdclient.LxdClient{}
	worker.CronClient.Init()
	worker.initialize()
	go worker.startCronJob()

	a.Router.HandleFunc("/api/v1/containers", handler.GetContainersHandler).Methods("GET")
	a.Router.HandleFunc("/api/v1/container/{name}", handler.GetContainerHandler).Methods("GET")
	a.Router.HandleFunc("/api/v1/container", handler.CreateContainerHandler).Methods("POST")
	a.Router.HandleFunc("/api/v1/container", handler.DeleteContainerHandler).Methods("DELETE")
	a.Router.HandleFunc("/api/v1/container/updatestate", handler.UpdateStateContainerHandler).Methods("POST")
}
