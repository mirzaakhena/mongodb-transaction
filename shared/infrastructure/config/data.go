package config

type Config struct {
	Server   Server   `json:"server"`
	Database Database `json:"database"`
}

type Server struct {
	Port int `json:"port,omitempty"`
}

type Database struct {
	URI string `json:"uri"`
}
