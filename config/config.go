package config

import (
	_ "embed"
	"os"

	"gopkg.in/yaml.v3"
)

//go:embed config.yml
var ConfigFile string

var Config struct {
	Rules []struct {
		Name  string `yaml:"name"`
		Regex string `yaml:"regex"`
		Bump  string `yaml:"bump"`
	} `yaml:"rules"`
}

func init() {
	if _, err := os.Stat("./.semver.yml"); err == nil {
		ConfigFile, err := os.ReadFile("./.semver.yml")
		if err != nil {
			panic(err)
		}

		if err := yaml.Unmarshal(ConfigFile, &Config); err != nil {
			panic(err)
		}
		return
	}

	if err := yaml.Unmarshal([]byte(ConfigFile), &Config); err != nil {
		panic(err)
	}
}
