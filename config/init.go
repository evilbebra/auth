package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	ApiApp struct {
		Port string `yaml:"port"`
	} `yaml:"api-app"`

	AuthApp struct {
		Port string `yaml:"port"`
	} `yaml:"auth-app"`

	Auth struct {
		SigningKey string `yaml:"signing_key"`
		TokenTTL   int    `yaml:"token_ttl"`
	} `yaml:"auth"`
}

var (
	c    *Config
	once sync.Once
)

func Init(filepath string) *Config {
	once.Do(func() {
		c = &Config{}
		log.Println("Read configuration..")
		if err := cleanenv.ReadConfig(filepath, c); err != nil {
			help, _ := cleanenv.GetDescription(c, nil)
			log.Println(help)
			log.Fatal(err)
		}
	})
	return c
}
