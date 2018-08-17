package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	lxdclient "github.com/go-squads/saga-agent/lxdclient"
	"github.com/jasonlvhit/gocron"
	"github.com/lxc/lxd/shared/api"
	log "github.com/sirupsen/logrus"
)

// CronWorker ...
type CronWorker struct {
	CronClient lxdclient.Client
	Cron       *gocron.Scheduler
}

type lxc struct {
	ID          string `db:"id" json:"id"`
	LxdID       string `db:"lxd_id" json:"lxd_id"`
	Name        string `db:"name" json:"name"`
	Type        string `db:"type" json:"type"`
	Alias       string `db:"alias" json:"alias"`
	Protocol    string `db:"protocol" json:"protocol"`
	Server      string `db:"server" json:"server"`
	Address     string `db:"address" json:"address"`
	Status      string `db:"status" json:"status"`
	Description string `db:"description" json:"description"`
}

func (cw *CronWorker) initialize() {
	cw.Cron = gocron.NewScheduler()
}

func (cw *CronWorker) startCronJob() {
	cw.Cron.Every(5).Seconds().Do(cw.doCron)
	<-cw.Cron.Start()
}

func (cw *CronWorker) doCron() {
	log.Infof("-- Cron Job Running every 5 seconds, sync to LXD : %s --", os.Getenv("LXD_NAME"))

	url := fmt.Sprintf(os.Getenv("SCHEDULER_URL") + "/lxd/" + os.Getenv("LXD_NAME") + "/lxc")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Infof(err.Error())
	}

	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(req)

	if err != nil {
		log.Infof(err.Error())
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Infof(err.Error())
	}

	lxcList := []lxc{}

	err = json.Unmarshal(body, &lxcList)

	if err != nil {
		log.Infof(err.Error())
	}

	cw.syncLxcStatus(lxcList)
}

func (cw *CronWorker) syncLxcStatus(lxcList []lxc) {
	for _, v := range lxcList {
		switch lxcStatus := v.Status; lxcStatus {
		case "creating":
			cw.createNewLxcSync(v)
		case "starting":
		case "stopping":
		case "deleting":
		default:
			log.Infof("lxc : %s, already synchronized", v.Name)
		}
	}
}

func (cw *CronWorker) createNewLxcSync(newLxc lxc) {
	request := api.ContainersPost{
		Name: newLxc.Name,
		Source: api.ContainerSource{
			Type:     newLxc.Type,
			Protocol: newLxc.Protocol,
			Server:   newLxc.Server,
			Alias:    newLxc.Alias,
		},
	}

	op, err := cw.CronClient.CreateContainer(request)
	if err != nil {
		log.Infof(err.Error())
	}

	err = op.Wait()
	if err != nil {
		log.Infof(err.Error())
	}

	log.Infof("Finish creating new container : %s", newLxc.Name)

	newLxc.Status = "stopped"
	cw.requestUpdateLxcStatus(newLxc)
}

func (cw *CronWorker) requestUpdateLxcStatus(updateLxcData lxc) {
	url := fmt.Sprintf(os.Getenv("SCHEDULER_URL") + "/lxc/" + updateLxcData.ID)
	payload, err := json.Marshal(updateLxcData)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))

	if err != nil {
		log.Infof(err.Error())
	}
	client := &http.Client{Timeout: 10 * time.Second}
	_, err = client.Do(req)

	if err != nil {
		log.Infof(err.Error())
	}

	log.Infof("Success updating lxc %s status to %s", updateLxcData.Name, updateLxcData.Status)
}
