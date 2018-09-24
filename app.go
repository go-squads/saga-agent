package main

import (
	"github.com/go-squads/saga-agent/lxdclient"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
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
	err := godotenv.Load()
	if err != nil {
		log.Panic("Load env failed")
		panic(err)
	}
	worker := CronWorker{}
	worker.CronClient = &lxdclient.LxdClient{}
	worker.CronClient.Init()
	worker.initialize()
	go worker.startCronJob()
}
