package util

import (
	"log"
	"main/src/config"
	"math/rand"
)

var (
	samplePlayers []config.SamplePlayer = nil
)

// GetOnlinePlayerCount returns the amount of online players using values calculated from the config.
func GetOnlinePlayerCount(conf *config.Config) (online int) {
	if conf.Players.Online.Random {
		online = rand.Intn(conf.Players.Online.Max-conf.Players.Online.Min) + conf.Players.Online.Min
	} else {
		online = conf.Players.Online.Value
	}

	return
}

// GetMaxPlayerCount returns the maximum numbers of players using values calculated from the config.
func GetMaxPlayerCount(conf *config.Config) (max int) {
	if conf.Players.Max.Random {
		max = rand.Intn(conf.Players.Max.Max-conf.Players.Max.Min) + conf.Players.Max.Min
	} else {
		max = conf.Players.Max.Value
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
func AddSamplePlayer(uuid string) error {
	profile, err := lookupProfile(uuid)

	if err != nil {
		return err
	}

	for i, player := range samplePlayers {
		if player.UUID != uuid {
			continue
		}

		samplePlayers = append(samplePlayers[0:i], samplePlayers[i+1:]...)

		break
	}

	samplePlayers = append([]config.SamplePlayer{
		{
			Username: profile.Username,
			UUID:     profile.UUID,
		},
	}, samplePlayers...)

	log.Printf("Added sample player %s (%s)\n", profile.Username, profile.UUID)

	return nil
}
