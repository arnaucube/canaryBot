package main

import (
	"fmt"

	"github.com/fatih/color"
)

const version = "0.0.1-dev"

func main() {
	savelog()

	color.Green("canaryBot")
	color.Yellow("version: " + version)

	color.Magenta("------matrix------")
	readConfig()
	loginMatrix()
	fmt.Print("matrix token: ")
	color.Cyan(matrixToken.AccessToken)
	fmt.Println("")
	matrixSendMsg("canaryBot started")

	color.Magenta("------starting to listen------")

	checker(config.Services)
}
