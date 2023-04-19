package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type minecraftProfile struct {
	UUID     string `json:"id"`
	Username string `json:"name"`
}

func lookupProfile(uuid string) (*minecraftProfile, error) {
	resp, err := http.Get(fmt.Sprintf("https://sessionserver.mojang.com/session/minecraft/profile/%s", url.PathEscape(strings.ReplaceAll(uuid, "-", ""))))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("mojang: unexpected status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result minecraftProfile

	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
