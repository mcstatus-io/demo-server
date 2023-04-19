package bedrock

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"time"
)

var (
	magic []byte = []byte{0x00, 0xff, 0xff, 0x00, 0xfe, 0xfe, 0xfe, 0xfe, 0xfd, 0xfd, 0xfd, 0xfd, 0x12, 0x34, 0x56, 0x78}
)

func readUnconnectedPingPacket(r io.Reader) error {
	// Packet ID - byte
	{
		var packetID byte

		if err := binary.Read(r, binary.BigEndian, &packetID); err != nil {
			return err
		}

		if packetID != 0x01 {
			return fmt.Errorf("bedrock: unexpected packet type: 0x%02X", packetID)
		}
	}

	// Time - long
	{
		var time int64

		if err := binary.Read(r, binary.BigEndian, &time); err != nil {
			return err
		}
	}

	// Magic - [16]byte
	{
		data := make([]byte, 16)

		if _, err := r.Read(data); err != nil {
			return err
		}

		if !bytes.Equal(data, magic) {
			return fmt.Errorf("bedrock: received magic data that does not match expected data: %X", data)
		}
	}

	// Client GUID - long
	{
		var clientGUID int64

		if err := binary.Read(r, binary.BigEndian, &clientGUID); err != nil {
			return err
		}
	}

	return nil
}

func writeUnconnectedPongPacket(w io.Writer) error {
	// Packet ID - byte
	if err := binary.Write(w, binary.BigEndian, byte(0x1C)); err != nil {
		return err
	}

	// Time - long
	if err := binary.Write(w, binary.BigEndian, time.Now().UnixMilli()); err != nil {
		return err
	}

	// Server GUID - long
	if err := binary.Write(w, binary.BigEndian, rand.Int63()); err != nil {
		return err
	}

	// Magic - [16]byte
	if _, err := w.Write(magic); err != nil {
		return err
	}

	// Server ID - string
	if err := writeString(w, getServerID()); err != nil {
		return err
	}

	return nil
}
