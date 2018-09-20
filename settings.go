package nec

type Command struct {
	Description string   `json:"description,omitempty"`
	Workers     int      `json:"workers"`
	OnErr       string   `json:"onerror"`
	Name        string   `json:"name"`
	Args        []string `json:"args,omitempty"`
}

type Settings struct {
	Commands map[string]*Command `json:"commands"`
	Paths    map[string][]string `json:"paths"`
}
