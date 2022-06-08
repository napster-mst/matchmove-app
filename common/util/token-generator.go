package util

import (
	ra "crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

func TokenGenerator() string {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(12-4) + 4
	b := make([]byte, x)
	ra.Read(b)
	return fmt.Sprintf("%x", b)
}
