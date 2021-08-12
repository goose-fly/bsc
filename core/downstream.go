package core

type DownstreamConfig struct {
	URI           string `toml:",omitempty"`
	Exchange      string `toml:",omitempty"`
	RoutingKey    string `toml:",omitempty"`
	RetryInterval int    `toml:",omitempty"` //in millsecond
}
