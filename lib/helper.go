package lib

import (
	"math/rand"
	"net/url"
)

func GenerateShortKey() string {
	// length, err := strconv.Atoi(os.Getenv("LENGTH"))
	// if err != nil {
	// 	log.Fatal("Error : ", err.Error())
	// 	return err.Error()
	// }

	// if length == 0 {
	// 	length = 6
	// }

	length := 6

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	shortKey := make([]byte, length)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}

	return string(shortKey)
}

func ValidateURL(URL string) bool {
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		return false
	}

	return true
}
