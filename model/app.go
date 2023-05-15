package model

type DBConfig struct {
	Name     string `env:"DBNAME"`
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	Username string `env:"USER"`
	Password string `env:"PASSWORD"`
}
