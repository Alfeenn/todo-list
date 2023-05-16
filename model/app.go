package model

type DBConfig struct {
	Name     string `env:"MYSQL_DBNAME"`
	Host     string `env:"MYSQL_HOST"`
	Port     string `env:"MYSQL_PORT"`
	Username string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD"`
}
