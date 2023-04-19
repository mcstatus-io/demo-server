package util

import "main/src/config"

// GetBedrockGamemodeName returns the capitalized name of the Bedrock Edition gamemode.
func GetBedrockGamemodeName(conf *config.Config) string {
	switch conf.BedrockEdition.Options.Gamemode {
	case "survival":
		return "Survival"
	case "creative":
		return "Creative"
	case "adventure":
		return "Adventure"
	case "spectator":
		return "Spectator"
	default:
		return "Unknown"
	}
}

// GetBedrockGamemodeID returns the gamemode ID of the Bedrock Edition gamemode.
func GetBedrockGamemodeID(conf *config.Config) int {
	switch conf.BedrockEdition.Options.Gamemode {
	case "survival":
		return 0
	case "creative":
		return 1
	case "adventure":
		return 2
	default:
		return 0
	}
}
