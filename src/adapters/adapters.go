package adapters

import (
	"net/url"
)

type Adapters struct {
	Adapters map[string]Adapter
}

func (adapters *Adapters) GetAdapter(monitorUrl string) Adapter {
	parsedUrl, urlErr := url.Parse(monitorUrl)
	if urlErr != nil {
		panic(urlErr)
	}
	return adapters.Adapters[parsedUrl.Hostname()]
}
