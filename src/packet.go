package main

import (
	"bufio"
	"bytes"
	"io"
)

func ReadPacket(r *bufio.ReadWriter) (int32, *bytes.Buffer, error) {
	packetLength, _, err := ReadVarInt(r)

	if err != nil {
		return 0, nil, err
	}

	data := make([]byte, packetLength)

	if _, err := r.Read(data); err != nil {
		return 0, nil, err
	}

	buf := bytes.NewBuffer(data)

	packetType, _, err := ReadVarInt(buf)

	if err != nil {
		return 0, nil, err
	}

	return packetType, buf, nil
}

func WritePacket(dst io.Writer, packetType int32, buf *bytes.Buffer) error {
	packetData := &bytes.Buffer{}

	if _, err := WriteVarInt(packetType, packetData); err != nil {
		return err
	}

	if _, err := io.Copy(packetData, buf); err != nil {
		return err
	}

	if _, err := WriteVarInt(int32(packetData.Len()), dst); err != nil {
		return err
	}

	if _, err := io.Copy(dst, packetData); err != nil {
		return err
	}

	return nil
}
