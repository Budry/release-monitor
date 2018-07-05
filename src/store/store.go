package store

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type Store map[string]time.Time

const storeFile = "/var/lib/release-monitor/data.json"

func InitializeStore() {
	if _, err := os.Stat(storeFile); os.IsNotExist(err) {
		fileErr := ioutil.WriteFile(storeFile, []byte("{}"), 0666)
		if fileErr != nil {
			panic(fileErr)
		}
	}
}

func UpdateStore(store Store) {
	marshaled, jsonErr := json.Marshal(store)
	if jsonErr != nil {
		panic(jsonErr)
	}
	fileErr := ioutil.WriteFile(storeFile, marshaled, 0644)
	if fileErr != nil {
		panic(fileErr)
	}
}

func GetStore() *Store {
	file, fileErr := os.Open(storeFile)
	if fileErr != nil {
		panic(fileErr)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var store *Store
	jsonErr := decoder.Decode(&store)
	if jsonErr != nil {
		panic(jsonErr)
	}

	return store
}
