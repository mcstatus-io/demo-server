package query

import (
	"encoding/binary"
	"fmt"
	"io"
)

func readBasePacket(r io.Reader) (byte, int32, error) {
	// Magic - unsigned short
	{
		var magic uint16

		if err := binary.Read(r, binary.BigEndian, &magic); err != nil {
			return 0, 0, err
		}

		if magic != 0xFEFD {
			return 0, 0, fmt.Errorf("query: received magic value with incorrect data: 0x%02X", magic)
		}
	}

	var packetType byte

	// Type - byte
	{
		if err := binary.Read(r, binary.BigEndian, &packetType); err != nil {
			return 0, 0, err
		}
	}

	var sessionID int32

	// Session ID - int32
	{
		if err := binary.Read(r, binary.BigEndian, &sessionID); err != nil {
			return 0, 0, err
		}

		sessionID &= 0x0F0F0F0F
	}

	return packetType, sessionID, nil
}
