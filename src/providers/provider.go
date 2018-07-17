package providers

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"sort"
	"strings"
	"sync"

	"github.com/budry/release-monitor/src/adapters"
	"github.com/budry/release-monitor/src/errors"
	"github.com/budry/release-monitor/src/monitors"
	"github.com/budry/release-monitor/src/releases"
	"github.com/budry/release-monitor/src/store"
)

type Provider struct {
	Adapters adapters.Adapters
}

func (provider *Provider) ProcessMonitor(monitor monitors.Monitor, wg *sync.WaitGroup) {
	defer wg.Done()

	adapter := provider.Adapters.GetAdapter(monitor.Adapter)
	if adapter == nil {
		panic("Missing adapter for monitor '" + monitor.Name + "'")
	}
	releaseRecords := adapter.GetReleases(&monitor)
	if len(releaseRecords) > 0 {

		sort.Sort(releases.ReleasesByDate(releaseRecords))

		storedData := *store.GetStore()
		_, exist := storedData[monitor.Name]

		var newReleases []releases.ReleaseRecord
		if !exist {
			storedData[monitor.Name] = releaseRecords[0].Date
			newReleases = append(newReleases, releaseRecords[0])
		}

		for _, release := range releaseRecords {
			if storedData[monitor.Name].Before(release.Date) {
				storedData[monitor.Name] = release.Date
				newReleases = append(newReleases, release)
			}
		}
		store.UpdateStore(storedData)

		if len(newReleases) > 0 {
			for _, release := range newReleases {
				for _, command := range monitor.Commands {
					newReleaseString, jsonErr := json.Marshal(release)
					errors.HandleError(jsonErr)

					command = strings.Replace(command, "%%RELEASE%%", string(newReleaseString), -1)
					out, err := exec.Command("sh", "-c", command).Output()
					errors.HandleError(err)
					fmt.Printf("Command: %s", command)
					fmt.Printf(" | Result: %s\n", out)
				}
			}
		}
	}
}

func (provider *Provider) Process(monitorItems []monitors.Monitor) {
	wg := &sync.WaitGroup{}
	for _, monitor := range monitorItems {
		wg.Add(1)
		go provider.ProcessMonitor(monitor, wg)
	}
	wg.Wait()
}
