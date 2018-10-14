package main

import (
	"encoding/json"
	"io/ioutil"
)

//MatrixConfig stores the matrix data config
type MatrixConfig struct {
	RoomId   string `json:"room_id"`
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
}
type Service struct {
	Name       string `json:"name"`
	Url        string `json:"url"`
	StatusCode int    `json:"statusCode"`
	Counter    int
}
type Config struct {
	Matrix    MatrixConfig `json:"matrix"`
	Services  []Service    `json:"services"`
	SleepTime int          `json:"sleepTime"`
	Retry     int          `json:"retry"`
}

var config Config

func readConfig() {
	file, e := ioutil.ReadFile("config.json")
	if e != nil {
		panic(e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &config)
}
