package config

import (
	"log"
	"os"
	"strconv"
)

func NewConfig() *Config {
	return &Config{
		Host: getValue("HOST"),
		Port: getIntValue("PORT"),
	}
}

func getValue(key string) string {
	return os.Getenv(key)
}

func getIntValue(key string) int16 {
	value, err := strconv.ParseInt(getValue(key), 10, 16)

	if err != nil {
		log.Fatal("Error parsing " + key)
	}

	return int16(value)
}
