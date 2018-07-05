package github

import (
	"strings"
	"time"

	"github.com/budry/release-monitor/src/errors"
	"github.com/budry/release-monitor/src/monitors"
	"github.com/budry/release-monitor/src/releases"
	"github.com/mmcdole/gofeed"
)

type GithubAdapter struct{}

func (github *GithubAdapter) GetReleases(monitor *monitors.Monitor) []releases.ReleaseRecord {
	fp := gofeed.NewParser()
	feed, feedError := fp.ParseURL(monitor.Url + "/releases.atom")
	errors.HandleError(feedError);

	var versions []releases.ReleaseRecord
	for _, release := range feed.Items {
		parsedID := strings.Split(release.GUID, "/")
		parsedTime, parsedTimeError := time.Parse(time.RFC3339, release.Updated)
		errors.HandleError(parsedTimeError)
		versions = append(versions, releases.ReleaseRecord{Date: parsedTime, Tag: parsedID[len(parsedID)-1]})
	}

	return versions
}
