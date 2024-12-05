package config

type Config struct {
	GatewayConfig GatewayConfig `json:"gateway"`
}

type GatewayConfig struct {
	ListenPort     int    `json:"listen_port"`
	ShortenerHost  string `json:"shortener_host"`
	ShortenertPort int    `json:"shortener_port"`
}
