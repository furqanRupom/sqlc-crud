package util

import (
	"log"
	"os"
	"strconv"
)

func CheckEnvStr(envValue string) string {
	val := os.Getenv(envValue)
	if val == "" {
		log.Printf("%s is required !", envValue)
	}
	return val
}

func CheckEnvInt(envValue string) int {
	val := os.Getenv(envValue)
	if val == "" {
		log.Printf("%s is required !", envValue)
	}
	valInt, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Printf("%s is not valid int!", envValue)
	}
	return int(valInt)
}

func CheckEnvBool(envValue string) bool {
	val := os.Getenv(envValue)
	if val == "" {
		log.Printf("%s is required", envValue)
	}
	valBool, err := strconv.ParseBool(val)
	if err != nil {
		log.Printf("%s is not valid bool!", envValue)
	}
	return valBool
}
