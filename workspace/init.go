package workspace

import (
	"fmt"
	"jinyaoma/go-experiment/config"
	"os"
	"path"
)

func InitUserWorkspace(userAccount string) error {
	targetPath := path.Join(config.WORKSPACE, userAccount)
	err := os.Mkdir(targetPath, os.ModeDir)
	return err
}

func GetInfo() {
	if info, err := os.Stat(config.WORKSPACE); err == nil {
		//size := reflect.ValueOf(info.Sys()).Elem().FieldByName("Size").Field(0).Int()
		fmt.Printf("%+v", info.Sys())
	}
}
