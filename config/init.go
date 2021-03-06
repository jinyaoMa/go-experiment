package config

import (
	"math/rand"
	"time"
)

const (
	WORKSPACE                 string = "D:/WORKSPACE"
	WORKSPACE_AVAILABLE_SPACE uint64 = 280 * 1024 * 1024 * 1024 // 280 GB
	WORKSPACE_CAN_INITIALIZE  bool   = true

	SERVER_PORT string = ":55699"

	DEFAULT_SIGNUP_ROLE_NAME string = "member"

	PASSWORD_DERECT_ACCESS_TO_ADMIN_WOKSPACE string = "jinyaoMa"
)

var (
	User_Limit int64 = 5
)

var (
	randSource rand.Source
)

func init() {
	randSource = rand.NewSource(time.Now().Unix())
}

func GenerateToken(len int) (string, time.Time) {
	t := time.Now().AddDate(0, 1, 0) // add a month
	r := rand.New(rand.NewSource(t.Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		var b int
		temp := r.Intn(10)
		if temp%2 == 0 { // letters
			if len%2 == 0 { // uppercase
				b = r.Intn(26) + 65
			} else { // lowercase
				b = r.Intn(26) + 97
			}
		} else { // digits
			b = r.Intn(10) + 48
		}
		bytes[i] = byte(b)
	}
	return string(bytes), t
}

func GenerateShareCode(len int) string {
	r := rand.New(randSource)
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		var b int
		temp := r.Intn(10)
		if temp%2 == 0 { // letters
			b = r.Intn(26) + 97
		} else { // digits
			b = r.Intn(10) + 48
		}
		bytes[i] = byte(b)
	}
	return string(bytes)
}
