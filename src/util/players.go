package util

import (
	"main/src/config"
	"math/rand"

	"github.com/google/uuid"
)

var (
	samplePlayers []config.SamplePlayer = nil
)

// GetJavaOnlinePlayerCount returns the amount of online players using values calculated from the config.
func GetJavaOnlinePlayerCount(conf *config.Config) (online int) {
	if conf.JavaEdition.Options.Players.Online.Random {
		online = rand.Intn(conf.JavaEdition.Options.Players.Online.Max-conf.JavaEdition.Options.Players.Online.Min) + conf.JavaEdition.Options.Players.Online.Min
	} else {
		online = conf.JavaEdition.Options.Players.Online.Value
	}

	return
}

// GetJavaMaxPlayerCount returns the maximum numbers of players using values calculated from the config.
func GetJavaMaxPlayerCount(conf *config.Config) (max int) {
	if conf.JavaEdition.Options.Players.Max.Random {
		max = rand.Intn(conf.JavaEdition.Options.Players.Max.Max-conf.JavaEdition.Options.Players.Max.Min) + conf.JavaEdition.Options.Players.Max.Min
	} else {
		max = conf.JavaEdition.Options.Players.Max.Value
	}

	return
}

// GetBedrockOnlinePlayerCount returns the amount of online players using values calculated from the config.
func GetBedrockOnlinePlayerCount(conf *config.Config) (online int) {
	if conf.BedrockEdition.Options.Players.Online.Random {
		online = rand.Intn(conf.BedrockEdition.Options.Players.Online.Max-conf.BedrockEdition.Options.Players.Online.Min) + conf.BedrockEdition.Options.Players.Online.Min
	} else {
		online = conf.BedrockEdition.Options.Players.Online.Value
	}

	return
}

// GetBedrockMaxPlayerCount returns the maximum numbers of players using values calculated from the config.
func GetBedrockMaxPlayerCount(conf *config.Config) (max int) {
	if conf.BedrockEdition.Options.Players.Max.Random {
		max = rand.Intn(conf.BedrockEdition.Options.Players.Max.Max-conf.BedrockEdition.Options.Players.Max.Min) + conf.BedrockEdition.Options.Players.Max.Min
	} else {
		max = conf.BedrockEdition.Options.Players.Max.Value
	}

	return
}

// GetSamplePlayers returns the sample player information about the server.
func GetSamplePlayers() []config.SamplePlayer {
	return samplePlayers
}

// SetSamplePlayers is the initial function called to set the sample players from the config file.
func SetSamplePlayers(players []config.SamplePlayer) {
	samplePlayers = players
}

// AddSamplePlayer is called if `config.java_status.add_logins_to_sample_players` is set, and it adds the player attempting to connect to the server to the sample players list.
func AddSamplePlayer(playerUUID string) error {
	profile, err := lookupProfile(playerUUID)

	if err != nil {
		return err
	}

	u, err := uuid.Parse(profile.UUID)

	if err != nil {
		return err
	}

	for i, player := range samplePlayers {
		if player.UUID != u.String() {
			continue
		}

		samplePlayers = append(samplePlayers[0:i], samplePlayers[i+1:]...)

		break
	}

	samplePlayers = append([]config.SamplePlayer{
		{
			Username: profile.Username,
			UUID:     u.String(),
		},
	}, samplePlayers...)

	return nil
}
