package util

import "os"

// GetEnv returns app envorinment : e.g. development, production, staging, testing, etc
func GetEnv() string {
	return os.Getenv("APP_ENV")
}

// IsProductionEnv returns whether the app is running using production env
func IsProductionEnv() bool {
	return os.Getenv("APP_ENV") == "production"
}

// IsDevelopmentEnv returns whether the app is running using production env
func IsDevelopmentEnv() bool {
	return os.Getenv("APP_ENV") == "development"
}

func IsFileorDirExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
