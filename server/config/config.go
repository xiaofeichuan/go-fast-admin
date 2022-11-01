package config

type Config struct {
	Jwt      JwtConfig      `yaml:"jwt"`
	Database DatabaseConfig `yaml:"database"`
}

type JwtConfig struct {
	Issuer string `yaml:"issuer"`
}
type DatabaseConfig struct {
	Default string `yaml:"default"`
}
