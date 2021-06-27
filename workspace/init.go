package workspace

import (
	"fmt"
	"jinyaoma/go-experiment/config"
	"os"
	"path/filepath"
)

type UserFilesFunc func(path string, isDir bool)

func InitUserWorkspace(userAccount string, fn UserFilesFunc) error {
	targetPath := filepath.Join(config.WORKSPACE, userAccount)
	err := os.Mkdir(targetPath, os.ModeDir)
	if os.IsExist(err) {
		initUserFiles(targetPath, fn)
		return nil
	}
	return err
}

func GetInfo() {
	disk := InitDisk(config.WORKSPACE)
	fmt.Printf("%+v", disk)
}

func initUserFiles(userPath string, fn UserFilesFunc) {
	err := filepath.Walk(userPath, func(p string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		fn(p, f.IsDir())
		return nil
	})
	if err != nil {
		fmt.Printf("initUserFiles:filepath.Walk() returned %v\n", err)
	}
}
