package config

type HostConfig struct {
	Host       string
	Port       int
	AuthHeader string
}

type Config struct {
	Host     HostConfig
	Interval int
}
