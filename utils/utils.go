package utils

import (
	"strconv"
	"time"
)

func GenerateRandomNumber() string {
	now := time.Now().UTC().Unix()
	return strconv.Itoa(int(now))
}
