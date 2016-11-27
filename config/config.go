package config

type Config struct {
	Host       string `env:"TEAMCITY_HOST,required"`
	Port       int    `env:"TEAMCITY_PORT" envDefault:"8111"`
	ProjectID  string `env:"TEAMCITY_PROJECT_ID,required"`
	AuthHeader string `env:"TEAMCITY_AUTH_HEADER,required"`
	Interval   int    `env:"TEAMCITY_UPDATE_INTERVAL" envDefault:"15"`
}
