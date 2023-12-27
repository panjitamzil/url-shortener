package lib

import (
	"log"
	"math/rand"
	"os"
	"strconv"
)

func GenerateShortKey() string {
	length, err := strconv.Atoi(os.Getenv("LENGTH"))
	if err != nil {
		log.Fatal("Error : ", err.Error())
		return err.Error()
	}

	if length == 0 {
		length = 6
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	shortKey := make([]byte, length)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}

	return string(shortKey)
}
