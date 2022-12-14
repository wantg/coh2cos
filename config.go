package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

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
    minWidth: 1300
    minHeight: 1000
`

var Config *config

func initConfig() {
	executablePath, _ := os.Executable()
	executableDirPath := filepath.Dir(executablePath)
	configFilePath := filepath.Join(executableDirPath, configFileName)

	if _, err := os.Stat(configFilePath); errors.Is(err, os.ErrNotExist) {
		os.WriteFile(configFilePath, []byte(strings.TrimSpace(defaultConfigMatadata)), os.ModePerm)
	}
	bts, _ := os.ReadFile(configFilePath)
	Config = &config{}
	yaml.Unmarshal(bts, Config)
	Config.ExeDirPath = filepath.Join(executableDirPath)
	Config.LogPath = parseSysVar(Config.LogPath)
}

func parseSysVar(s string) string {
	parsedVar := s
	sysVarList := []string{"USERPROFILE"}
	for _, v := range sysVarList {
		_v := "%" + v + "%"
		if strings.Contains(parsedVar, _v) {
			parsedVar = strings.ReplaceAll(parsedVar, _v, os.Getenv(v))
		}
	}
	return parsedVar
}
