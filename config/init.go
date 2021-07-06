package config

import (
	"math/rand"
	"time"
)

const (
	WORKSPACE                 string = "D:/WORKSPACE"
	WORKSPACE_AVAILABLE_SPACE uint64 = 32 * 1024 * 1024 * 1024 // 32 GB
	WORKSPACE_CAN_INITIALIZE  bool   = true

	SERVER_PORT string = ":3000"

	DEFAULT_SIGNUP_ROLE_NAME string = "member"
)

var (
	User_Limit uint = 5
)

func GenerateToken(len int) (string, time.Time) {
	t := time.Now().AddDate(0, 1, 0) // add a month
	r := rand.New(rand.NewSource(t.Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes), t
}

func GenerateShareCode(start int, end int) int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(end - start)
	random = start + random
	return random
}
