package main

import (
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/model"
)

func main() {
	model.InitDB(config.WORK_SPACE)
}
