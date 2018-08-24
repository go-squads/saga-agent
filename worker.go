package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

// Interface for testing
type HttpClient interface {
	Do(*http.Request) *http.Response
	NewRequest(string, string, io.Reader) (*http.Request, error)
}

type IOHandler interface {
	ReadAll(io.Reader) ([]byte, error)
}

// Ends here

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

	if len(lxcList) > 0 {

		ip, err := cw.extractLxcIPAddress(lxcList[0])
		if err != nil {
			log.Error(err.Error())
		}
		log.Infof("IP Address extracted: %s for lxc %s", ip, lxcList[0].Name)
	}

	// cw.syncLxcStatus(lxcList)

}

func (cw *CronWorker) syncLxcStatus(lxcList []lxc) {
	for _, v := range lxcList {
		switch lxcStatus := v.Status; lxcStatus {
		case "creating":
			cw.createNewLxcSync(v)
		case "starting":
			v.Status = "start"
			cw.updateLxcStateSync(v)
		case "stopping":
			v.Status = "stop"
			cw.updateLxcStateSync(v)
		case "deleting":
			cw.deleteLxcSync(v)
		default:
			log.Infof("lxc : %s, already synchronized", v.Name)
		}
	}
}

func (cw *CronWorker) createNewLxcSync(newLxcData lxc) {
	request := api.ContainersPost{
		Name: newLxcData.Name,
		Source: api.ContainerSource{
			Type:     newLxcData.Type,
			Protocol: newLxcData.Protocol,
			Server:   newLxcData.Server,
			Alias:    newLxcData.Alias,
		},
	}

	// Add another state: Failed to create
	op, err := cw.CronClient.CreateContainer(request)
	if err != nil {
		log.Error("Container can't be cretaed")
		newLxcData.Status = "Failed to create"
		cw.requestUpdateLxcStatus(newLxcData)
	}

	if err = op.Wait(); err != nil {
		log.Errorf("Container creation operation failed")
		newLxcData.Status = "Failed to create"
		cw.requestUpdateLxcStatus(newLxcData)
		return
	}

	log.Infof("Finish creating new container : %s", newLxcData.Name)

	newLxcData.Status = "stopped"
	cw.requestUpdateLxcStatus(newLxcData)
}

func (cw *CronWorker) updateLxcStateSync(updateLxcData lxc) {
	request := api.ContainerStatePut{
		Action:  updateLxcData.Status,
		Timeout: 60,
	}

	op, err := cw.CronClient.UpdateContainerState(updateLxcData.Name, request)

	if err != nil {
		log.Infof(err.Error())
	}

	if err = op.Wait(); err != nil {
		log.Infof(err.Error())
	}

	updateLxcData.Status = cw.changeLxcStateString(updateLxcData.Status)
	log.Infof("Container %s state is now : %s", updateLxcData.Name, updateLxcData.Status)
	cw.requestUpdateLxcStatus(updateLxcData)
}

func (cw *CronWorker) changeLxcStateString(currentState string) string {
	if currentState == "start" {
		return "started"
	}
	return "stopped"
}

func (cw *CronWorker) deleteLxcSync(deleteLxcData lxc) {
	op, err := cw.CronClient.DeleteContainer(deleteLxcData.Name)

	if err != nil {
		log.Infof(err.Error())
	}

	if err = op.Wait(); err != nil {
		log.Infof(err.Error())
	}

	log.Infof("Finish deleting container : %s", deleteLxcData.Name)
	cw.requestDeleteLxc(deleteLxcData)
}

func (cw *CronWorker) requestUpdateLxcStatus(updateLxcData lxc) {
	url := fmt.Sprintf(os.Getenv("SCHEDULER_URL") + "/lxc")
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

func (cw *CronWorker) requestDeleteLxc(deleteLxcData lxc) {
	url := fmt.Sprintf(os.Getenv("SCHEDULER_URL") + "/lxc")
	payload, err := json.Marshal(deleteLxcData)
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(payload))

	if err != nil {
		log.Infof(err.Error())
	}
	client := &http.Client{Timeout: 10 * time.Second}
	_, err = client.Do(req)

	if err != nil {
		log.Infof(err.Error())
	}

	log.Infof("Success deleting lxc %s from database", deleteLxcData.Name)
}

func (cw *CronWorker) extractLxcIPAddress(lxcToCheck lxc) (string, error) {
	state, _, err := cw.CronClient.GetContainerState(lxcToCheck.Name)
	if err != nil {
		return "", err
	}
	return state.Network["eth0"].Addresses[0].Address, nil
}
