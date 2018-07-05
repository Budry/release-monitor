package store

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/budry/release-monitor/src/errors"
)

type Store map[string]time.Time

const storeFile = "/var/lib/release-monitor/data.json"

func InitializeStore() {
	if _, err := os.Stat(storeFile); os.IsNotExist(err) {
		fileErr := ioutil.WriteFile(storeFile, []byte("{}"), 0666)
		errors.HandleError(fileErr)
	}
}

func UpdateStore(store Store) {
	marshaled, jsonErr := json.Marshal(store)
	errors.HandleError(jsonErr)
	fileErr := ioutil.WriteFile(storeFile, marshaled, 0666)
	errors.HandleError(fileErr)
}

func GetStore() *Store {
	file, fileErr := os.Open(storeFile)
	errors.HandleError(fileErr)
	defer file.Close()

	decoder := json.NewDecoder(file)
	var store *Store
	jsonErr := decoder.Decode(&store)
	errors.HandleError(jsonErr)

	return store
}
