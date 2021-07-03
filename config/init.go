package config

import (
	"math/rand"
	"time"
)

const (
	WORKSPACE                 = "D:/WORKSPACE"
	WORKSPACE_AVAILABLE_SPACE = 32 * 1024 * 1024 * 1024 // 32 GB
	WORKSPACE_CAN_INITIALIZE  = true

	SERVER_PORT = ":3000"
)

func GenerateToken(len int) (string, time.Time) {
	t := time.Now()
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
