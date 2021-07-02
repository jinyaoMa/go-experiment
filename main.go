package main

import (
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/controller"
	"jinyaoma/go-experiment/model"
)

func main() {
	model.InitDB(config.WORKSPACE)

	controller.RunRouter()
}
