package main

import (
	"github.com/go-squads/saga-agent/lxdclient"
)

// App ...
type App struct{}

// Run ...
func (a *App) Run() {
	for true {

	}
}

// Initialize ...
func (a *App) Initialize() {
	worker := CronWorker{}
	worker.CronClient = &lxdclient.LxdClient{}
	worker.CronClient.Init()
	worker.initialize()
	go worker.startCronJob()
}
