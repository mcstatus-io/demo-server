package config

import (
	"os"
	"time"

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
		Enable                  bool          `yaml:"enable"`
		Host                    string        `yaml:"host"`
		Port                    uint16        `yaml:"port"`
		GlobalSessionExpiration time.Duration `yaml:"global_session_expiration"`
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
	Mods struct {
		Enable     bool `yaml:"enable"`
		FMLVersion int  `yaml:"fml_version"`
		Mods       []struct {
			ID      string `yaml:"id"`
			Version string `yaml:"version"`
		} `yaml:"mods"`
	} `yaml:"mods"`
	Software string `yaml:"software"`
	Plugins  []struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"plugins"`
	MapName string `yaml:"map_name"`
}

// ReadFile reads the YAML file from the path.
func (c *Config) ReadFile(file string) error {
	data, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, c)
}
