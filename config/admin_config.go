package config

type CoreConfiguration struct {
	Version string
}

type AdminConfiguration struct {
	Core     CoreConfiguration
	Uuid     string
	Hostname string
	Address  string
}
