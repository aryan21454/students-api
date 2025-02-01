package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string
}

type Config struct {
	Env          string `yaml:"env"`
	Storage_path string `yaml:"storage_path"`
	HTTPServer   `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if  configPath == "" {
		flags := flag.String("config", "", "path to config file")
	
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			log.Fatal("config path is required")
		}
	}

	if _ , err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", configPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("failed to read config file: %s", err)
	}
	return &cfg

}