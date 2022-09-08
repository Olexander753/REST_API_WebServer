package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Listen struct {
		Port string `yaml:"port" env-default:"8080"`
		IP   string `yaml:"ip" env-default:"127.0.0.1"`
	} `yaml:"listen"`
	MongoDB struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		Database   string `yaml:"database"`
		AuthDB     string `yaml:"authdb"`
		Collection string `yaml:"collection"`
	} `yaml:"mongodb"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	//выполниться ровно один раз, при следующих вызовах просто будет возвращать instance
	once.Do(func() {
		log.Println("Read configuration")
		instance = &Config{}
		err := cleanenv.ReadConfig("config.yaml", instance)
		if err != nil {
			dis, _ := cleanenv.GetDescription(instance, nil)
			log.Println(dis)
			log.Fatal(err)

		}
	})
	return instance
}
