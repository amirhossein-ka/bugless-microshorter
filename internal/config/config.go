package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	GatewayConfig GatewayConfig `json:"gateway"`
}

type GatewayConfig struct {
	ListenPort    int    `json:"listen_port"`
	ShortenerHost string `json:"shortener_host"`
	ShortenerPort int    `json:"shortener_port"`
	CacheSize     int    `json:"cache_size"`
}

func Parse(cfg *Config, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(cfg)
}
