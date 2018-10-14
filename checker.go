package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func checker(services []Service) {
	log.Println("serverChecker started")
	log.Println(services)

	ticker := time.NewTicker(time.Second * time.Duration(int64(config.SleepTime)))
	for _ = range ticker.C {
		for k, service := range services {
			resp, err := checkUrl(service.Url)
			if err != nil {
				if service.Counter == 0 {
					msg := "⚠️ Service " + service.Name + ", with url " + service.Url + " is not responding"
					matrixSendMsg(msg)
				}
				services[k].Counter = services[k].Counter + 1
			} else if resp.StatusCode != service.StatusCode {
				if service.Counter == 0 {
					msg := "⚠️ Service " + service.Name + ", with url " + service.Url + " has returned a non expected StatusCode: " + strconv.Itoa(resp.StatusCode) + ", expected: " + strconv.Itoa(service.StatusCode)
					matrixSendMsg(msg)
				}
				services[k].Counter = services[k].Counter + 1
			} else {
				if services[k].Counter != 0 {
					msg := "✔️ Service " + service.Name + ", with url " + service.Url + " is alive again"
					matrixSendMsg(msg)
				}
				services[k].Counter = 0
			}

			if services[k].Counter > config.Retry {
				msg := "⚠️ Service " + services[k].Name + ", with url " + services[k].Url + " is still failing"
				matrixSendNotice(msg)
				services[k].Counter = 1
			}
		}
	}
}

func checkUrl(url string) (*http.Response, error) {
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	return resp, err
}
