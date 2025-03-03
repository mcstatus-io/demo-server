package java

import (
	"main/src/config"

	"main/src/util"
)

type status struct {
	Version     version     `json:"version"`
	Players     players     `json:"players"`
	Description config.Chat `json:"description"`
	Favicon     string      `json:"favicon,omitempty"`
	ModInfo     *modInfo    `json:"modinfo,omitempty"`
	ForgeData   *forgeData  `json:"forgeData,omitempty"`
}

type version struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type players struct {
	Online int                   `json:"online"`
	Max    int                   `json:"max"`
	Sample []config.SamplePlayer `json:"sample"`
}

type samplePlayer struct {
	Username string `json:"name"`
	UUID     string `json:"id"`
}

type modInfo struct {
	Type    string      `json:"type"`
	ModList []legacyMod `json:"modList"`
}

type legacyMod struct {
	ID      string `json:"modid"`
	Version string `json:"version"`
}

type forgeData struct {
	Channels          []any      `json:"channels"`
	Mods              []forgeMod `json:"mods"`
	FMLNetworkVersion int        `json:"fmlNetworkVersion"`
}

type forgeMod struct {
	ID      string `json:"modId"`
	Version string `json:"modmarker"`
}

func getStatusResponse() (result status) {
	result = status{
		Version: version{
			Name:     conf.JavaEdition.Options.Version.Name,
			Protocol: conf.JavaEdition.Options.Version.Protocol,
		},
		Players: players{
			Online: util.GetJavaOnlinePlayerCount(conf),
			Max:    util.GetJavaMaxPlayerCount(conf),
			Sample: util.GetSamplePlayers(),
		},
		Description: conf.JavaEdition.Options.MOTD,
		Favicon:     favicon,
	}

	if conf.JavaEdition.Options.Mods.Enable {
		switch conf.JavaEdition.Options.Mods.FMLVersion {
		case 1:
			{
				result.ModInfo = &modInfo{
					Type:    "FML",
					ModList: make([]legacyMod, 0),
				}

				for _, mod := range conf.JavaEdition.Options.Mods.List {
					result.ModInfo.ModList = append(result.ModInfo.ModList, legacyMod{
						ID:      mod.ID,
						Version: mod.Version,
					})
				}

				break
			}
		case 2:
			{
				result.ForgeData = &forgeData{
					FMLNetworkVersion: 2,
					Channels:          make([]any, 0),
					Mods:              make([]forgeMod, 0),
				}

				for _, mod := range conf.JavaEdition.Options.Mods.List {
					result.ForgeData.Mods = append(result.ForgeData.Mods, forgeMod{
						ID:      mod.ID,
						Version: mod.Version,
					})
				}

				break
			}
		}
	}

	return
}
