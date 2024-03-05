package utils

import "os"

const (
	AppVersionEnv  = "APP_VERSION"
	AppVersionKey  = "version"
	ServiceNameEnv = "SERVICE_NAME"
	ServiceNameKey = "service"
)

func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
