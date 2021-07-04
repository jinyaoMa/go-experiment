package workspace

import (
	"errors"
	"fmt"
	"jinyaoma/go-experiment/config"
	"os"
	"path/filepath"
)

const (
	WORKSPACE_ERROR_WORKSPACE_EXIST = "workspace exist"
)

var (
	disk Disk
)

func init() {
	var err error
	disk, err = InitDisk(config.WORKSPACE)
	if err != nil {
		fmt.Println(err)
	}
}

func GetDisk() *Disk {
	return &disk
}

type UserFilesFunc func(path string, fileInofo os.FileInfo)

func NewUserWorkspace(userAccount string, fn UserFilesFunc) error {
	targetPath := filepath.Join(config.WORKSPACE, userAccount)
	err := os.Mkdir(targetPath, os.ModeDir)
	if os.IsExist(err) {
		return errors.New(WORKSPACE_ERROR_WORKSPACE_EXIST)
	}
	initUserFiles(targetPath, fn)
	return nil
}

func InitUserWorkspace(userAccount string, fn UserFilesFunc) error {
	if !config.WORKSPACE_CAN_INITIALIZE {
		return nil
	}
	targetPath := filepath.Join(config.WORKSPACE, userAccount)
	err := os.Mkdir(targetPath, os.ModeDir)
	if os.IsExist(err) {
		initUserFiles(targetPath, fn)
		return nil
	}
	return err
}

func initUserFiles(userPath string, fn UserFilesFunc) {
	err := filepath.Walk(userPath, func(p string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		fn(p, f)
		return nil
	})
	if err != nil {
		fmt.Printf("initUserFiles:filepath.Walk() returned %v\n", err)
	}
}
