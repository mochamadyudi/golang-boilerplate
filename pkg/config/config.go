package config

type Config struct {
	Server   Server
	Database Database
	Jwt      JsonWebToken
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Name     string
	TimeZone string
	SSlMode  string
}

type JsonWebToken struct {
	Secret      string
	ExpiresHour string
}
