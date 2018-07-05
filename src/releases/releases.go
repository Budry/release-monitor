package releases

import "time"

type ReleaseRecord struct {
	Date time.Time `json:"date"`
	Tag  string    `json:"tag"`
}

type ReleasesByDate []ReleaseRecord

func (releases ReleasesByDate) Len() int           { return len(releases) }
func (releases ReleasesByDate) Swap(i, j int)      { releases[i], releases[j] = releases[j], releases[i] }
func (releases ReleasesByDate) Less(i, j int) bool { return releases[i].Date.Before(releases[j].Date) }
