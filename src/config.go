package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config is the configuration data for this application.
type Config struct {
	JavaStatus struct {
		Enable         bool   `yaml:"enable"`
		Host           string `yaml:"host"`
		Port           uint16 `yaml:"port"`
		LogConnections bool   `yaml:"log_connections"`
	} `yaml:"java_status"`
	BedrockStatus struct {
		Enable         bool   `yaml:"enable"`
		Host           string `yaml:"host"`
		Port           uint16 `yaml:"port"`
		LogConnections bool   `yaml:"log_connections"`
	} `yaml:"bedrock_status"`
	Query struct {
		Enable         bool   `yaml:"enable"`
		Host           string `yaml:"host"`
		Port           uint16 `yaml:"port"`
		LogConnections bool   `yaml:"log_connections"`
	} `yaml:"query"`
	Players struct {
		Online struct {
			Random bool `yaml:"random"`
			Value  int  `yaml:"value"`
			Min    int  `yaml:"min"`
			Max    int  `yaml:"max"`
		} `yaml:"online"`
		Max struct {
			Random bool `yaml:"random"`
			Value  int  `yaml:"value"`
			Min    int  `yaml:"min"`
			Max    int  `yaml:"max"`
		} `yaml:"max"`
		Sample []string `yaml:"sample"`
	} `yaml:"players"`
	Version struct {
		Name     string `yaml:"name"`
		Protocol int    `yaml:"protocol"`
	} `yaml:"version"`
	MOTD string `yaml:"motd"`
}

// ReadFile reads the YAML file from the path.
func (c *Config) ReadFile(file string) error {
	data, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, c)
}
