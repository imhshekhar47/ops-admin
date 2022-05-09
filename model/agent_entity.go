package model

type Location struct {
	Latitude  string
	Longitude string
}

type AgentInfo struct {
	Group       string
	Component   string
	Environment string
	Site        string
	Location    Location
}

type AgentRefEntity struct {
	AgentId string
	Key     string
	Value   string

	Info AgentInfo
}
