package config

type Config struct {
}

func (conf *Config) PortNumber() uint16 {
	return 8080
}
