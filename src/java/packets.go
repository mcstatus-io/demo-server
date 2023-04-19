package java

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/src/config"
)

const (
	nextStateStatus int32 = iota + 1
	nextStateLogin
)

// [C -> S] Handshake (0x00)
func readHandshakePacket(r io.Reader) (int32, error) {
	packetType, data, err := readPacket(r)

	if err != nil {
		return 0, err
	}

	if packetType != 0x00 {
		return 0, fmt.Errorf("handshake: unexpected packet type: 0x%02X", packetType)
	}

	// 1. Protocol Version - varint
	{
		if _, err = readVarInt(data); err != nil {
			return 0, err
		}
	}

	// 2. Server Address - string
	{
		if _, err = readString(data); err != nil {
			return 0, err
		}
	}

	// 3. Server Port - unsigned short
	{
		var port uint16

		if err := binary.Read(data, binary.BigEndian, &port); err != nil {
			return 0, err
		}
	}

	var nextState int32

	// 4. Next State - varint
	{
		nextState, err = readVarInt(data)

		if err != nil {
			return 0, err
		}

		if nextState != 1 && nextState != 2 {
			return 0, fmt.Errorf("handshake: unexpected next state value: %d", nextState)
		}
	}

	if data.Len() > 0 {
		return 0, fmt.Errorf("handshake: %d bytes left over at end of packet", data.Len())
	}

	return nextState, nil
}

// [C -> S] Request (0x00)
func readRequestPacket(r io.Reader) error {
	packetType, data, err := readPacket(r)

	if err != nil {
		return err
	}

	if packetType != 0x00 {
		return fmt.Errorf("request: unexpected packet type: 0x%02X", packetType)
	}

	if data.Len() > 0 {
		return fmt.Errorf("request: %d bytes left over at end of packet", data.Len())
	}

	return nil
}

// [C <- S] Response (0x00)
func writeResponsePacket(w io.Writer) error {
	data := &bytes.Buffer{}

	response, err := json.Marshal(getStatusResponse())

	if err != nil {
		log.Fatal(err)
	}

	// 1. JSON Response
	{
		if err = writeString(data, string(response)); err != nil {
			return err
		}
	}

	if err = writePacket(w, 0x00, data); err != nil {
		return err
	}

	return nil
}

// [C -> S] Ping (0x01)
func readPingPacket(r io.Reader) (int64, error) {
	packetType, data, err := readPacket(r)

	if err != nil {
		return 0, err
	}

	if packetType != 0x01 {
		return 0, fmt.Errorf("ping: unexpected packet type: 0x%02X", packetType)
	}

	var payload int64

	// 1. Payload - long
	{
		if err = binary.Read(data, binary.BigEndian, &payload); err != nil {
			return 0, err
		}
	}

	if data.Len() > 0 {
		return 0, fmt.Errorf("ping: %d bytes left over at end of packet", data.Len())
	}

	return payload, nil
}

// [C <- S] Pong (0x01)
func writePongPacket(w io.Writer, payload int64) error {
	data := &bytes.Buffer{}

	// 1. Payload - long
	if err := binary.Write(data, binary.BigEndian, payload); err != nil {
		return err
	}

	if err := writePacket(w, 0x01, data); err != nil {
		return err
	}

	return nil
}

// [C -> S] Login Start (0x00)
func readLoginStartPacket(r io.Reader) (*string, error) {
	packetType, data, err := readPacket(r)

	if err != nil {
		return nil, err
	}

	if packetType != 0x00 {
		return nil, fmt.Errorf("login start: unexpected packet type: 0x%02X", packetType)
	}

	// 1. Name - string
	{
		if _, err = readString(data); err != nil {
			return nil, err
		}
	}

	var hasUUID bool

	// 2. Has Player UUID - boolean
	{
		uuid, err := data.ReadByte()

		if err != nil {
			return nil, err
		}

		hasUUID = uuid == 1
	}

	var uuid *string = nil

	// 3. Optional UUID
	{
		if hasUUID {
			uuidData := make([]byte, 16)

			if _, err := data.Read(uuidData); err != nil {
				return nil, err
			}

			value := hex.EncodeToString(uuidData)
			uuid = &value
		}
	}

	if data.Len() > 0 {
		return nil, fmt.Errorf("login start: %d bytes left over at end of packet", data.Len())
	}

	return uuid, nil
}

// [C <- S] Pong (0x01)
func writeDisconnectPacket(w io.Writer, reason config.Chat) error {
	data := &bytes.Buffer{}

	// 1. Reason - Chat
	{
		reasonData, err := json.Marshal(reason)

		if err != nil {
			return err
		}

		if err := writeString(data, string(reasonData)); err != nil {
			return err
		}
	}

	if err := writePacket(w, 0x00, data); err != nil {
		return err
	}

	return nil
}
