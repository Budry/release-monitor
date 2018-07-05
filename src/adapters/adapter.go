package adapters

import (
	"bitbucket.org/budry/release-monitor/src/monitors"
	"bitbucket.org/budry/release-monitor/src/releases"
)

type Adapter interface {
	GetReleases(*monitors.Monitor) []releases.ReleaseRecord
}
