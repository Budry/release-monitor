package monitors

type Monitor struct {
	Name     string   `json:"name"`
	Url      string   `json:"url"`
	Commands []string `json:"commands"`
	Adapter  string   `json:"adapter"`
}
