package config

import "log"

type ConfigMode int

const (
	PRODUCTION ConfigMode = iota
	DEVELOPMENT
)

var configModes = [...]string{
	"production",
	"development",
}

func (mode ConfigMode) String() string {
	return configModes[(mode)]
}

type DdlAutoMode int

const (
	CREATE DdlAutoMode = iota
	NONE
)

var ddlModes = [...]string{
	"CREATE",
	"NONE",
}

var ddlModeMap = map[string]DdlAutoMode{
	"CREATE": CREATE,
	"NONE":   NONE,
}

func (mode DdlAutoMode) String() string {
	return ddlModes[mode]
}

func MapToDdlAuto(mode string) DdlAutoMode {
	autoMode, ok := ddlModeMap[mode]
	if !ok {
		log.Default().Println("Not valid DDL_AUTO mode. Initializing in NONE mode")
		autoMode = NONE
	}
	return autoMode
}
