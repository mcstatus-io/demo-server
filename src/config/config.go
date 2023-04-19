package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config is the configuration data for this application.
type Config struct {
	JavaEdition struct {
		Status struct {
			Enable                        bool   `yaml:"enable"`
			Host                          string `yaml:"host"`
			Port                          uint16 `yaml:"port"`
			DisconnectReason              Chat   `yaml:"disconnect_reason"`
			AddLoginsToSamplePlayers      bool   `yaml:"add_logins_to_sample_players"`
			AddLoginsToSamplePlayersLimit int    `yaml:"add_logins_to_sample_players_limit"`
		} `yaml:"status"`
		Query struct {
			Enable                  bool          `yaml:"enable"`
			Host                    string        `yaml:"host"`
			Port                    uint16        `yaml:"port"`
			GlobalSessionExpiration time.Duration `yaml:"global_session_expiration"`
		} `yaml:"query"`
		Options struct {
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
				Sample []SamplePlayer `yaml:"sample"`
			} `yaml:"players"`
			Version struct {
				Name     string `yaml:"name"`
				Protocol int    `yaml:"protocol"`
			} `yaml:"version"`
			MOTD Chat `yaml:"motd"`
			Mods struct {
				Enable     bool `yaml:"enable"`
				FMLVersion int  `yaml:"fml_version"`
				List       []struct {
					ID      string `yaml:"id"`
					Version string `yaml:"version"`
				} `yaml:"list"`
			} `yaml:"mods"`
			Software string `yaml:"software"`
			Plugins  []struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
			} `yaml:"plugins"`
			MapName string `yaml:"map_name"`
		} `yaml:"options"`
	} `yaml:"java_edition"`
	BedrockEdition struct {
		Status struct {
			Enable bool   `yaml:"enable"`
			Host   string `yaml:"host"`
			Port   uint16 `yaml:"port"`
		} `yaml:"status"`
		Options struct {
			Edition string `yaml:"edition"`
			MOTD    struct {
				Line1 string `yaml:"line_1"`
				Line2 string `yaml:"line_2"`
			} `yaml:"motd"`
			Version struct {
				Name     string `yaml:"name"`
				Protocol int    `yaml:"protocol"`
			} `yaml:"version"`
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
			} `yaml:"players"`
			Gamemode string `yaml:"gamemode"`
		} `yaml:"options"`
	} `yaml:"bedrock_edition"`
}

// ReadFile reads the YAML file from the path.
func (c *Config) ReadFile(file string) error {
	data, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, c)
}
