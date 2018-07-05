package adapters

import (
	"net/url"

	"github.com/budry/release-monitor/src/errors"
)

type Adapters struct {
	Adapters map[string]Adapter
}

func (adapters *Adapters) GetAdapter(monitorUrl string) Adapter {
	parsedUrl, urlErr := url.Parse(monitorUrl)
	errors.HandleError(urlErr)

	return adapters.Adapters[parsedUrl.Hostname()]
}
