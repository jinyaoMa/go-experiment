package model

type User struct {
	Username string
	Account  string
	Password string
	Type     string
	Active   string
}

func NewUser() *User {
	return &User{
		Username: "",
		Account:  "",
		Password: "",
		Type:     UT_GUEST,
		Active:   UA_INACTIVE,
	}
}

func (user *User) ToString() string {
	return "Username: " + user.Username + "\n" +
		"Account: " + user.Account + "\n" +
		"Password: " + user.Password + "\n" +
		"Type: " + user.Type + "\n" +
		"Active: " + user.Active + "\n\n"
}
