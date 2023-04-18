package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config is the configuration data for this application.
type Config struct {
	JavaStatus struct {
		Enable           bool   `yaml:"enable"`
		Host             string `yaml:"host"`
		Port             uint16 `yaml:"port"`
		DisconnectReason Chat   `yaml:"disconnect_reason"`
	} `yaml:"java_status"`
	BedrockStatus struct {
		Enable bool   `yaml:"enable"`
		Host   string `yaml:"host"`
		Port   uint16 `yaml:"port"`
	} `yaml:"bedrock_status"`
	Query struct {
		Enable bool   `yaml:"enable"`
		Host   string `yaml:"host"`
		Port   uint16 `yaml:"port"`
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
		Sample []struct {
			Username string `yaml:"username"`
			UUID     string `yaml:"uuid"`
		} `yaml:"sample"`
	} `yaml:"players"`
	Version struct {
		Name     string `yaml:"name"`
		Protocol int    `yaml:"protocol"`
	} `yaml:"version"`
	MOTD Chat `yaml:"motd"`
}

// ReadFile reads the YAML file from the path.
func (c *Config) ReadFile(file string) error {
	data, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, c)
}
