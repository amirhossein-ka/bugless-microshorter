package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	GatewayConfig   GatewayConfig   `json:"gateway"`
	ShortenerConfig ShortenerConfig `json:"shortener"`
}

type GatewayConfig struct {
	ListenPort    int    `json:"listen_port"`
	ShortenerHost string `json:"shortener_host"`
	ShortenerPort int    `json:"shortener_port"`
	CacheSize     int    `json:"cache_size"`
}

type ShortenerConfig struct {
	ListenPort int `json:"listen_port"`

	DBHost     string `json:"db_host"`
	DBPort     int    `json:"db_port"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
}

func Parse(cfg *Config, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(cfg)
}
