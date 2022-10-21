package main

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type config struct {
	ExeDirPath string `json:"ExeDirPath" yaml:"exeDirPath"`
	LogPath    string `json:"LogPath"    yaml:"logPath"`
	Window     struct {
		MinWidth  int `json:"MinWidth"   yaml:"minWidth"`
		MinHeight int `json:"MinHeight"  yaml:"minHeight"`
	} `json:"Window" yaml:"window"`
}

const configMatadata = `
logPath: '%USERPROFILE%\Documents\My Games\Company of Heroes 2\warnings.log'
window:
    minWidth: 1280
    minHeight: 1000
`

var Config *config

func initConfig() {
	Config = &config{}
	yaml.Unmarshal([]byte(configMatadata), Config)
	executablePath, _ := os.Executable()
	Config.ExeDirPath = filepath.Join(filepath.Dir(executablePath))
}
