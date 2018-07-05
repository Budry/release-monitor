package adapters

import (
	"github.com/budry/release-monitor/src/monitors"
	"github.com/budry/release-monitor/src/releases"
)

type Adapter interface {
	GetReleases(*monitors.Monitor) []releases.ReleaseRecord
}
