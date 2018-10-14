package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

//MatrixToken stores the token data from matrix
type MatrixToken struct {
	AccessToken string `json:"access_token"`
	Server      string `json:"server"`
	UserId      string `json:"user_id"`
	DeviceId    string `json:"device_id"`
}

var matrixToken MatrixToken

var defaultTimeFormat = "15:04:05"

func loginMatrix() {
	url := config.Matrix.Server + "/_matrix/client/r0/login"
	jsonStr := `{
		"type":"m.login.password",
		"user":"` + config.Matrix.User + `",
		"password":"` + config.Matrix.Password + `"
	}`
	b := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("POST", url, b)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal([]byte(body), &matrixToken)

}
func matrixSendMsg(msg string) {
	txnId := strconv.Itoa(rand.Int())
	url := config.Matrix.Server + "/_matrix/client/r0/rooms/" + config.Matrix.RoomId + "/send/m.room.message/" + txnId + "?access_token=" + matrixToken.AccessToken
	jsonStr := `{
		"body":"` + time.Now().Format(defaultTimeFormat) + `: ` + msg + `",
		"msgtype":"m.text"
	}`
	b := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("PUT", url, b)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	fmt.Print("msg sent to Matrix: ")
	color.Green(msg)

}

func matrixSendNotice(msg string) {
	txnId := strconv.Itoa(rand.Int())
	url := config.Matrix.Server + "/_matrix/client/r0/rooms/" + config.Matrix.RoomId + "/send/m.room.message/" + txnId + "?access_token=" + matrixToken.AccessToken
	jsonStr := `{
		"body":"` + time.Now().Format(defaultTimeFormat) + `: ` + msg + `",
		"msgtype":"m.notice"
	}`
	b := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("PUT", url, b)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	fmt.Print("msg sent to Matrix: ")
	color.Green(msg)

}
