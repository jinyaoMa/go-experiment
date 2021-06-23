package main

import (
	fs "jinyaoma/go-experiment/fs"
	"jinyaoma/go-experiment/model"
)

var (
	WORK_SPACE = "D:\\WORKSPACE"
)

func main() {
	userfile, err := fs.CreateUserFile(WORK_SPACE)
	if err != nil {
		panic(err)
	}

	user := model.NewUser()
	user.Username = "Jinyao Ma"
	user.Account = "admin"
	user.Password = "admin"
	user.Type = model.UT_ADMIN
	user.Active = model.UA_INACTIVE

	userfile.Insert(user)
	userfile.Insert(user)
}
