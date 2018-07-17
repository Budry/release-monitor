package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/budry/release-monitor/src/errors"
	"github.com/budry/release-monitor/src/monitors"
	"github.com/budry/release-monitor/src/releases"
)

type GitlabAdapter struct{}

type GitlabTag struct {
	Name   string          `json:"name"`
	Commit GitlabTagCommit `json:"commit"`
}

type GitlabTagCommit struct {
	Created string `json:"created_at"`
}

type GitlabTags []GitlabTag

func (github *GitlabAdapter) GetReleases(monitor *monitors.Monitor) []releases.ReleaseRecord {
	response, err := http.Get(monitor.Url)
	errors.HandleError(err)
	defer response.Body.Close()

	dec := json.NewDecoder(response.Body)
	tags := GitlabTags{}
	err = dec.Decode(&tags)
	fmt.Println(tags)

	var versions []releases.ReleaseRecord
	for _, tag := range tags {
		parsedTime, parsedTimeError := time.Parse(time.RFC3339, tag.Commit.Created)
		errors.HandleError(parsedTimeError)
		versions = append(versions, releases.ReleaseRecord{Date: parsedTime, Tag: tag.Name})
	}

	return versions
}
