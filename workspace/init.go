package workspace

import (
	"errors"
	"jinyaoma/go-experiment/config"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	WORKSPACE_ERROR_WORKSPACE_EXIST        = "workspace exist"
	WORKSPACE_ERROR_WORKSPACE_SPECIAL_CHAR = "account contains special characters"
)

var (
	disk Disk
)

func init() {
	var err error
	disk, err = InitDisk(config.WORKSPACE)
	if err != nil {
		log.Println(err)
	}
}

func GetDisk() *Disk {
	return &disk
}

type UserFilesFunc func(path string, fileInofo os.FileInfo)

func NewUserWorkspace(userAccount string, fn UserFilesFunc) error {
	if strings.ContainsAny(userAccount, "\\/:*?\"<>|") {
		return errors.New(WORKSPACE_ERROR_WORKSPACE_SPECIAL_CHAR)
	}
	targetPath := filepath.Join(config.WORKSPACE, userAccount)
	err := os.Mkdir(targetPath, os.ModeDir)
	if os.IsExist(err) {
		return errors.New(WORKSPACE_ERROR_WORKSPACE_EXIST)
	}
	if err != nil {
		return err
	}
	initUserFiles(targetPath, fn)
	return err
}

func InitUserWorkspace(userAccount string, fn UserFilesFunc) error {
	if !config.WORKSPACE_CAN_INITIALIZE {
		return nil
	}
	if strings.ContainsAny(userAccount, "\\/:*?\"<>|") {
		return errors.New(WORKSPACE_ERROR_WORKSPACE_SPECIAL_CHAR)
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
		relPath, errRelPath := filepath.Rel(userPath, p)
		if errRelPath != nil {
			return errRelPath
		}
		fn(relPath, f)
		return nil
	})
	if err != nil {
		log.Printf("initUserFiles:filepath.Walk() returned %v\n", err)
	}
}
