package config

import (
	"encoding/json"
	"os"

	"github.com/budry/release-monitor/src/errors"
	"github.com/budry/release-monitor/src/monitors"
)

type Configuration struct {
	Interval string `json:"interval"`
	Monitors []monitors.Monitor `json:"monitors"`
}

const configFile = "/etc/release-monitor/config.json"
var globalConfiguration *Configuration

func GetGlobalConfiguration() *Configuration {
	if globalConfiguration == nil {
		file, fileErr := os.Open(configFile)
		errors.HandleError(fileErr)
		defer file.Close()

		decoder := json.NewDecoder(file)
		jsonErr := decoder.Decode(&globalConfiguration)
		errors.HandleError(jsonErr)
	}

	return globalConfiguration
}
