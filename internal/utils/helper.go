package utils

import "os"

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != ""{
		return value
	}
	return defaultValue
}

func TotalPages(totalSize, pageSize int64) int64 {
	totalPages := (totalSize + pageSize - 1) / pageSize
	return totalPages
}