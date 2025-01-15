package models

type Settings struct {
	ConnectionString ConnectionString `json:"ConnectionString"`
	Development      Development      `json:"Development"`
	LaunchSettings   LaunchSettings   `json:"LaunchSettings"`
}

type ConnectionString struct {
	DefaultConnection string `json:"DefaultConnection"`
}

type Development struct {
	Log Log `json:"Log"`
}

type Log struct {
	Debug bool `json:"Debug"`
}

type LaunchSettings struct {
	Port   int `json:"Port"`
	Secure int `json:"Secure"`
}
