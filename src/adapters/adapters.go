package adapters

type Adapters struct {
	Adapters map[string]Adapter
}

func (adapters *Adapters) GetAdapter(adapterName string) Adapter {
	return adapters.Adapters[adapterName]
}
