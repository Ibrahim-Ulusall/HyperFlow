package models

type Settings struct {
	ConnectionString ConnectionString `json:"ConnectionString"`
}

type ConnectionString struct {
	DefaultConnection string `json:"DefaultConnection"`
}
