package workspace

import (
	"fmt"
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/model"
	"os"
	"path"
)

func InitWorkspace(users []model.User) {
	for _, user := range users {
		os.Mkdir(path.Join(config.WORKSPACE, user.Account), os.ModeDir)
	}
	fmt.Println("Workspace initialized")
}
