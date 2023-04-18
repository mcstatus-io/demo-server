package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
	"math/rand"
	"net"
)

type JavaStatus struct {
	Version     JavaStatusVersion `json:"version"`
	Players     JavaStatusPlayers `json:"players"`
	Description Chat              `json:"description"`
	Favicon     *string           `json:"favicon,omitempty"`
}

type JavaStatusVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type JavaStatusPlayers struct {
	Online int                      `json:"online"`
	Max    int                      `json:"max"`
	Sample []JavaStatusSamplePlayer `json:"sample,omitempty"`
}

type JavaStatusSamplePlayer struct {
	Username string `json:"name"`
	UUID     string `json:"id"`
}

type Chat struct {
	Text string `json:"text"`
}

func AcceptJavaStatusConnections() {
	for {
		conn, err := statusListener.Accept()

		if err != nil {
			continue
		}

		if config.JavaStatus.LogConnections {
			log.Printf("Received a status connection from %s\n", conn.RemoteAddr())
		}

		go HandleJavaStatusConnection(conn)
	}
}

func HandleJavaStatusConnection(conn net.Conn) {
	defer conn.Close()

	socket := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	// [C -> S] Handshake (0x00)
	{
		packetType, data, err := ReadPacket(socket)

		if err != nil {
			return
		}

		if packetType != 0x00 {
			return
		}

		// Protocol Version - varint
		if _, _, err = ReadVarInt(data); err != nil {
			return
		}

		// Server Address - string
		if _, err = ReadString(data); err != nil {
			return
		}

		// Server Port - unsigned short
		var port uint16

		if err := binary.Read(data, binary.BigEndian, &port); err != nil {
			return
		}

		// Next State - varint
		nextState, _, err := ReadVarInt(data)

		if err != nil {
			return
		}

		if nextState != 1 {
			return
		}

		if data.Len() > 0 {
			return
		}
	}

	// [C -> S] Request (0x00)
	{
		packetType, data, err := ReadPacket(socket)

		if err != nil {
			return
		}

		if packetType != 0x00 {
			return
		}

		if data.Len() > 0 {
			return
		}
	}

	// [C <- S] Response (0x00)
	{
		data := &bytes.Buffer{}

		response, err := json.Marshal(GetJavaStatusResponse())

		if err != nil {
			log.Fatal(err)
		}

		if err = WriteString(string(response), data); err != nil {
			return
		}

		if err = WritePacket(socket, 0x00, data); err != nil {
			return
		}

		if err = socket.Flush(); err != nil {
			return
		}
	}

	var payload int64

	// [C -> S] Ping (0x01)
	{
		packetType, data, err := ReadPacket(socket)

		if err != nil {
			return
		}

		if packetType != 0x01 {
			return
		}

		if err = binary.Read(data, binary.BigEndian, &payload); err != nil {
			return
		}
	}

	// [C <- S] Pong (0x01)
	{
		data := &bytes.Buffer{}

		if err := binary.Write(data, binary.BigEndian, payload); err != nil {
			return
		}

		if err := WritePacket(socket, 0x01, data); err != nil {
			return
		}

		if err := socket.Flush(); err != nil {
			return
		}
	}
}

func GetJavaStatusResponse() (result JavaStatus) {
	result = JavaStatus{
		Version: JavaStatusVersion{
			Name:     config.Version.Name,
			Protocol: config.Version.Protocol,
		},
		Players: JavaStatusPlayers{
			Sample: make([]JavaStatusSamplePlayer, 0),
		},
		Description: Chat{
			Text: config.MOTD,
		},
		Favicon: serverIcon,
	}

	if config.Players.Online.Random {
		result.Players.Online = rand.Intn(config.Players.Online.Max-config.Players.Online.Min) + config.Players.Online.Min
	} else {
		result.Players.Online = config.Players.Online.Value
	}

	if config.Players.Max.Random {
		result.Players.Max = rand.Intn(config.Players.Max.Max-config.Players.Max.Min) + config.Players.Max.Min
	} else {
		result.Players.Max = config.Players.Max.Value
	}

	return
}
