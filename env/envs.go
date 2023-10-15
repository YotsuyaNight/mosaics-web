package env

import "os"

func getEnvOr(name string, def string) string {
	env := os.Getenv(name)
	if env == "" {
		return def
	}
	return env
}

func GetBaseDir() string {
	return getEnvOr("BASE_DIR", ".")
}

func GetUploadsDir() string {
	return getEnvOr("UPLOADS_DIR", "uploads")
}

func GetResultDir() string {
	return getEnvOr("RESULT_DIR", "result")
}
