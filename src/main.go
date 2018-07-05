package main

import (
	"bitbucket.org/budry/release-monitor/src/adapters"
	"bitbucket.org/budry/release-monitor/src/adapters/github"
	"bitbucket.org/budry/release-monitor/src/config"
	"bitbucket.org/budry/release-monitor/src/providers"
	"bitbucket.org/budry/release-monitor/src/store"
)

func main() {
	store.InitializeStore()
	adaptersStruct := adapters.Adapters{
		Adapters: map[string]adapters.Adapter{
			"github.com": &github.GithubAdapter{},
		},
	}
	provider := &providers.Provider{Adapters: adaptersStruct}
	provider.Process(config.GetGlobalConfiguration().Monitors)
}
