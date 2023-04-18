package util

import (
	"main/src/config"
	"math/rand"
)

func GetOnlinePlayerCount(conf *config.Config) (online int) {
	if conf.Players.Online.Random {
		online = rand.Intn(conf.Players.Online.Max-conf.Players.Online.Min) + conf.Players.Online.Min
	} else {
		online = conf.Players.Online.Value
	}

	return
}

func GetMaxPlayerCount(conf *config.Config) (max int) {
	if conf.Players.Max.Random {
		max = rand.Intn(conf.Players.Max.Max-conf.Players.Max.Min) + conf.Players.Max.Min
	} else {
		max = conf.Players.Max.Value
	}

	return
}
