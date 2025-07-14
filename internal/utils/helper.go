package utils

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != ""{
		return value
	}
	return defaultValue
}

func GetIntEnv(key string, defaultValue int) int {
	value := os.Getenv(key);
	if value == ""{
		
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {

		return defaultValue
	}
	
	return intValue
}

func TotalPages(totalSize, pageSize int64) int64 {
	totalPages := (totalSize + pageSize - 1) / pageSize
	return totalPages
}

func RandomTimeSecond() time.Duration {
	min := 300
	max := 350
	randomNumber := rand.Intn(max - min + 1) + min

	return time.Duration(randomNumber) * time.Second
}