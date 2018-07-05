package config

import (
	"encoding/json"
	"os"

	"bitbucket.org/budry/release-monitor/src/monitors"
)

type Configuration struct {
	Monitors []monitors.Monitor `json:"monitors"`
}

const configFile = "/etc/release-monitor/config.json"
var globalConfiguration *Configuration

func GetGlobalConfiguration() *Configuration {
	if globalConfiguration == nil {
		file, fileErr := os.Open(configFile)
		if fileErr != nil {
			panic(fileErr)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		jsonErr := decoder.Decode(&globalConfiguration)
		if jsonErr != nil {
			panic(jsonErr)
		}
	}

	return globalConfiguration
}
