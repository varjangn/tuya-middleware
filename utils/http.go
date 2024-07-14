package utils

import "fmt"

func GetConfigPath(configPath string) string {
	if configPath != "" {
		return fmt.Sprintf("./config/config-%s", configPath)
	}
	return "./config/config-local"
}
