package entity

type Option struct {
	Message string `json:"message"`
	Value   string `json:"value"`
}

type Step struct {
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Optional bool     `json:"optional,omitempty"`
	Format   string   `json:"format,omitempty"`
	Message  string   `json:"message,omitempty"`
	Options  []Option `json:"options,omitempty"`
}

type Configuration struct {
	Name		string`json:"name"`
	Template	string`json:"template"`
	Steps 		[]Step `json:"steps"`
}