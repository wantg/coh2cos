package main

import (
	"errors"
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

const configFileName = "config.yaml"

const defaultConfigMatadata = `
logPath: '%USERPROFILE%\Documents\My Games\Company of Heroes 2\warnings.log'
window:
    minWidth: 1280
    minHeight: 1000
`

var Config *config

func initConfig() {
	executablePath, _ := os.Executable()
	executableDirPath := filepath.Dir(executablePath)
	configFilePath := filepath.Join(executableDirPath, configFileName)

	if _, err := os.Stat(configFilePath); errors.Is(err, os.ErrNotExist) {
		os.WriteFile(configFilePath, []byte(defaultConfigMatadata), os.ModePerm)
	}
	bts, _ := os.ReadFile(configFilePath)
	Config = &config{}
	yaml.Unmarshal(bts, Config)
	Config.ExeDirPath = filepath.Join(executableDirPath)
}
