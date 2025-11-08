package configs

import (
	"log"
	"os"
	"strconv"
)

func mustGetEnvString(key string) string {
	str := os.Getenv(key)
	if str == "" {
		log.Fatalln(key, "not found in env")
	}
	return str
}

func mustGetEnvInt(key string) int {
	str := mustGetEnvString(key)
	n, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalln(key, "invalid", str)
	}
	return n
}
