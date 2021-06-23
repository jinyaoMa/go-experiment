package fs

import (
	"errors"
	"jinyaoma/go-experiment/model"
	"os"
	"path"
)

type UserFile struct {
	*os.File
}

func CreateUserFile(directory string) (*UserFile, error) {
	var filename = path.Join(directory, "user"+FE_ADMIN_TXT)
	if isFileExist(filename) {
		err := errors.New(E_EXIST)
		return nil, err
	} else {
		userFile, err := os.Create(filename)
		return &UserFile{
			userFile,
		}, err
	}
}

func (userfile *UserFile) Insert(user *model.User) {
	userfile.WriteString(user.ToString())
}
