package configstore

type Group struct {
	ID      string    `json:"id"`
	Configs []Entries `json:"configs"`
	Version string    `json:"version"`
}

type Config struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	Entries Entries
}

type Entries struct {
	Entries map[string]string `json:"entries"`
}
