package config

type ConfigMode int

const (
	PRODUCTION ConfigMode = iota
	DEVELOPMENT
)

var configModes = [...]string{
	"production",
	"development",
}

func (mode ConfigMode) String() string { return configModes[(mode)] }
