package utils

import (
	"math/rand"
	"time"
)

// GetSerialCode 生产随机数
func GetSerialCode() int {

	rand.Seed(time.Now().UnixNano())
	var serialCode int
	for {
		serialCode = rand.Intn(1000000)
		if serialCode >= 111111 && serialCode <= 999999 {
			break
		}
	}
	return serialCode
}
