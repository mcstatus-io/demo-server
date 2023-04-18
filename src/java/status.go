package java

import (
	"main/src/config"
	"math/rand"
)

type status struct {
	Version     version     `json:"version"`
	Players     players     `json:"players"`
	Description config.Chat `json:"description"`
	Favicon     string      `json:"favicon,omitempty"`
}

type version struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type players struct {
	Online int            `json:"online"`
	Max    int            `json:"max"`
	Sample []samplePlayer `json:"sample"`
}

type samplePlayer struct {
	Username string `json:"name"`
	UUID     string `json:"id"`
}

func getStatusResponse() (result status) {
	result = status{
		Version: version{
			Name:     conf.Version.Name,
			Protocol: conf.Version.Protocol,
		},
		Players: players{
			Sample: make([]samplePlayer, 0),
		},
		Description: conf.MOTD,
		Favicon:     favicon,
	}

	if conf.Players.Online.Random {
		result.Players.Online = rand.Intn(conf.Players.Online.Max-conf.Players.Online.Min) + conf.Players.Online.Min
	} else {
		result.Players.Online = conf.Players.Online.Value
	}

	if conf.Players.Max.Random {
		result.Players.Max = rand.Intn(conf.Players.Max.Max-conf.Players.Max.Min) + conf.Players.Max.Min
	} else {
		result.Players.Max = conf.Players.Max.Value
	}

	for _, player := range conf.Players.Sample {
		result.Players.Sample = append(result.Players.Sample, samplePlayer{
			Username: player.Username,
			UUID:     player.UUID,
		})
	}

	return
}
