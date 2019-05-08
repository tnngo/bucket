package bucket

import (
	"math/rand"
	"time"
)

var seed = "0123456789abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func token(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	v := make([]uint8, 0)
	for i := 0; i < length && length > 0; i++ {
		rm := r.Int(71)
		v = append(v, seed[rm])
	}
	return string(v)
}
