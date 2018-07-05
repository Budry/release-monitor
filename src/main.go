package main

import (
	"fmt"
	"time"

	"github.com/budry/release-monitor/src/adapters"
	"github.com/budry/release-monitor/src/adapters/github"
	"github.com/budry/release-monitor/src/config"
	"github.com/budry/release-monitor/src/providers"
	"github.com/budry/release-monitor/src/store"
	"github.com/robfig/cron"
)

func run(provider *providers.Provider) {
	fmt.Print("Start processing at " + time.Now().String())
	provider.Process(config.GetGlobalConfiguration().Monitors)
	fmt.Println(" Done at " + time.Now().String())
}

func main() {
	store.InitializeStore()

	adaptersStruct := adapters.Adapters{
		Adapters: map[string]adapters.Adapter{
			"github.com": &github.GithubAdapter{},
		},
	}
	provider := &providers.Provider{Adapters: adaptersStruct}
	run(provider)

	c := cron.New()
	c.AddFunc(config.GetGlobalConfiguration().Interval, func() {
		run(provider)
	})
	c.Start()

	select {}
}
