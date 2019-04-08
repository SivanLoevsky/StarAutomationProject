package main

import (
	"crypto/tls"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var (
	panelIP = "172.16.101.54:3000"
)

func main(){
	success := pingPanel(panelIP)

	if !success {
		logrus.Errorln("PING NOT SUCCESSFUL... EXITING")
		os.Exit(1)
	}

	logrus.Println("END OF RUN... EXITING")
	os.Exit(0)
}

func pingPanel(panelAddr string) bool {
	logrus.Println("ATTEMPTING PING...")
	url := fmt.Sprintf("https://%s/api/v1/ping", panelAddr)

	httpc := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify:true,
			},
		},
	}

	pingResponse, err := httpc.Get(url)

	if err != nil {
		logrus.Errorln("PING ERR: ", err)
		return false
	} else if pingResponse.StatusCode != http.StatusOK {
		logrus.Errorln("PING BAD RESPONSE: ", pingResponse.StatusCode)
		return false
	} else {
		logrus.Println("PING SUCCESS")
		return true
	}
}