package providers

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"sort"
	"strings"

	"bitbucket.org/budry/release-monitor/src/adapters"
	"bitbucket.org/budry/release-monitor/src/errors"
	"bitbucket.org/budry/release-monitor/src/monitors"
	"bitbucket.org/budry/release-monitor/src/releases"
	"bitbucket.org/budry/release-monitor/src/store"
)

type Provider struct {
	Adapters adapters.Adapters
}

func (provider *Provider) Process(monitors []monitors.Monitor) {
	for _, monitor := range monitors {
		adapter := provider.Adapters.GetAdapter(monitor.Url)
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
				for _, command := range monitor.Commands {
					newReleasesString, jsonErr := json.Marshal(newReleases)
					errors.HandleError(jsonErr)

					command = strings.Replace(command, "%%RELEASE%%", string(newReleasesString), -1)
					out, err := exec.Command("sh", "-c", command).Output()
					errors.HandleError(err)
					fmt.Printf("Command: %s", command)
					fmt.Printf(" | Result: %s\n", out)
				}
			}
		}
	}
}
