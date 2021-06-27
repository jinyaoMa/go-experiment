package workspace

import (
	"jinyaoma/go-experiment/config"
	"os"
	"path"
)

func InitUserWorkspace(userAccount string) error {
	targetPath := path.Join(config.WORKSPACE, userAccount)
	err := os.Mkdir(targetPath, os.ModeDir)
	return err
}
