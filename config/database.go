package config

type Database struct {
	Host     string `env:"DB_HOST" envDefault:"mysql"`
	Port     string `env:"DB_PORT" envDefault:"3306"`
	User     string `env:"DB_USER" envDefault:"user"`
	Password string `env:"DB_PASSWORD" envDefault:"password"`
	Name     string `env:"DB_NAME" envDefault:"database"`
}
