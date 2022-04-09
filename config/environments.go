package config

import (
	"github.com/nicholasjackson/env"
)

func NewEnvironment() {
	_ = env.Parse()
}

var BindAddress = env.Int("BIND_ADDRESS", false, 8080, "Bind address for the server")
var LogLevel = env.String("LOG_LEVEL", false, "debug", "Log output level for the server [debug, info, trace]")
var buildMode = env.String("BUILD_MODE", false, "local", "Is build production server")

var localMode = "local"

func IsBuildMode() bool {
	return *buildMode != localMode
}
