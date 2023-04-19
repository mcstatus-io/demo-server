package bedrock

import (
	"fmt"
	"main/src/util"
	"math/rand"
)

// Edition (MCPE or MCEE for Education Edition);MOTD line 1;Protocol Version;Version Name;Player Count;Max Player Count;Server Unique ID;MOTD line 2;Game mode;Game mode (numeric);Port (IPv4);Port (IPv6);
func getServerID() string {
	return fmt.Sprintf(
		"%s;%s;%d;%s;%d;%d;%d;%s;%s;%d;%d;%d;",
		conf.BedrockEdition.Options.Edition,
		conf.BedrockEdition.Options.MOTD.Line1,
		conf.BedrockEdition.Options.Version.Protocol,
		conf.BedrockEdition.Options.Version.Name,
		util.GetBedrockOnlinePlayerCount(conf),
		util.GetBedrockMaxPlayerCount(conf),
		rand.Uint64(),
		conf.BedrockEdition.Options.MOTD.Line2,
		util.GetBedrockGamemodeName(conf),
		util.GetBedrockGamemodeID(conf),
		conf.BedrockEdition.Status.Port,
		conf.BedrockEdition.Status.Port,
	)
}
