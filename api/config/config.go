package config

import "github.com/caarlos0/env"

type Config struct {
	Env string `env:"ENV" envDefault:"dev"`
	Port string `env:"PORT" envDefault:"80"`
	DBHost string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBUser string `env:"DB_USER" envDefault:"todo"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"P@ssw0rd"`
	DBName string `env:"DB_NAME" envDefault:"todo"`
}

func New() (*Config,error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}