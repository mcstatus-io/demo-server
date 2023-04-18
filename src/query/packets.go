package query

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"main/src/util"
	"math/rand"
	"strconv"
	"strings"
)

func writeHandshakePacket(w io.Writer, sessionID int32) error {
	challengeToken := strconv.FormatInt(int64(rand.Int31()), 10)
	sessions[sessionID] = challengeToken

	// Type - byte
	if err := binary.Write(w, binary.BigEndian, byte(0x09)); err != nil {
		return err
	}

	// Session ID - int32
	if err := binary.Write(w, binary.BigEndian, sessionID); err != nil {
		return err
	}

	// Challenge Token - string
	if err := writeNTString(w, challengeToken); err != nil {
		return err
	}

	return nil
}

func readRequestPacket(r io.Reader, w io.Writer, sessionID int32) (bool, error) {
	if _, ok := sessions[sessionID]; !ok {
		return false, fmt.Errorf("query: invalid or expired session ID: %X", sessionID)
	}

	// Challenge Token - int32
	{
		var challengeToken int32

		if err := binary.Read(r, binary.BigEndian, &challengeToken); err != nil {
			return false, err
		}

		if sessions[sessionID] != strconv.FormatInt(int64(challengeToken), 10) {
			return false, fmt.Errorf("query: received challenge token did not match stored")
		}
	}

	var isFullStat bool

	// Padding - optional
	{
		data := make([]byte, 4)

		if _, err := r.Read(data); err != nil {
			if !errors.Is(err, io.EOF) {
				return false, err
			}

			isFullStat = false
		} else {
			isFullStat = true
		}
	}

	return isFullStat, nil
}

func writeBasicStatPacket(w io.Writer, sessionID int32) error {
	// Type - byte
	if err := binary.Write(w, binary.BigEndian, byte(0x00)); err != nil {
		return err
	}

	// Session ID - int32
	if err := binary.Write(w, binary.BigEndian, sessionID); err != nil {
		return err
	}

	// MOTD - null-terminated string
	if err := writeNTString(w, conf.MOTD.String()); err != nil {
		return err
	}

	// Game Type - null-terminated string
	if err := writeNTString(w, "SMP"); err != nil {
		return err
	}

	// Map Name - null-terminated string
	if err := writeNTString(w, conf.MapName); err != nil {
		return err
	}

	// Online Players - null-terminated string
	if err := writeNTString(w, strconv.FormatInt(int64(util.GetOnlinePlayerCount(conf)), 10)); err != nil {
		return err
	}

	// Max Players - null-terminated string
	if err := writeNTString(w, strconv.FormatInt(int64(util.GetMaxPlayerCount(conf)), 10)); err != nil {
		return err
	}

	// Host Port - little-endian short
	if err := binary.Write(w, binary.LittleEndian, conf.Query.Port); err != nil {
		return err
	}

	// Host IP - null-terminated string
	if err := writeNTString(w, conf.Query.Host); err != nil {
		return err
	}

	return nil
}

func writeFullStatPacket(w io.Writer, sessionID int32) error {
	// Type - byte
	if err := binary.Write(w, binary.BigEndian, byte(0x00)); err != nil {
		return err
	}

	// Session ID - int32
	if err := binary.Write(w, binary.BigEndian, sessionID); err != nil {
		return err
	}

	// Padding - [11]byte
	if _, err := w.Write([]byte{0x73, 0x70, 0x6C, 0x69, 0x74, 0x6E, 0x75, 0x6D, 0x00, 0x80, 0x00}); err != nil {
		return err
	}

	// K, V Section - null-terminated string pair
	{
		plugins := make([]string, 0)

		for _, plugin := range conf.Plugins {
			plugins = append(plugins, fmt.Sprintf("%s %s", plugin.Name, plugin.Version))
		}

		data := map[string]string{
			"hostname":   conf.MOTD.String(),
			"game_type":  "SMP",
			"game_id":    "MINECRAFT",
			"version":    conf.Version.Name,
			"plugins":    fmt.Sprintf("%s: %s", conf.Software, strings.Join(plugins, "; ")),
			"map":        conf.MapName,
			"numplayers": strconv.FormatInt(int64(util.GetOnlinePlayerCount(conf)), 10),
			"maxplayers": strconv.FormatInt(int64(util.GetMaxPlayerCount(conf)), 10),
			"hostport":   strconv.FormatUint(uint64(conf.Query.Port), 10),
			"hostip":     conf.Query.Host,
		}

		for key, value := range data {
			if err := writeNTString(w, key); err != nil {
				return err
			}

			if err := writeNTString(w, value); err != nil {
				return err
			}
		}

		if _, err := w.Write([]byte{0x00}); err != nil {
			return err
		}
	}

	// Padding - [10]byte
	if _, err := w.Write([]byte{0x01, 0x70, 0x6C, 0x61, 0x79, 0x65, 0x72, 0x5F, 0x00, 0x00}); err != nil {
		return err
	}

	// Players - null-terminated strings list
	{
		for _, player := range conf.Players.Sample {
			if err := writeNTString(w, player.Username); err != nil {
				return err
			}
		}

		if _, err := w.Write([]byte{0x00}); err != nil {
			return err
		}
	}

	return nil
}
