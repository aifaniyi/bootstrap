package gotemplate

// Settings :
const (
	Settings = `
	package settings

import (
	"gitlab.com/aifaniyi/go-libs/conf"
)

const (
	defaultPort = ":8000"
	defaultHost = "localhost"
)

// Settings : application configurations
type Settings struct {
	APIHost, APIPort string
	IsDev string
}

// LoadSettings : load application configuration
// from environement variables
func LoadSettings() *Settings {

	return &Settings{
		APIHost: conf.GetStringEnv("API_HOST", defaultHost),
		APIPort: conf.GetStringEnv("API_PORT", defaultPort),
		IsDev: conf.GetStringEnv("DEV", "true"),
	}
}

// Dependencies : struct containing application dependencies
// e.g postgres, cassandra, etc
type Dependencies struct {
	Db db.Service
}
`
)
