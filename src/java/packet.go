package java

import (
	"bytes"
	"io"
)

func readPacket(r io.Reader) (int32, *bytes.Buffer, error) {
	packetLength, err := readVarInt(r)

	if err != nil {
		return 0, nil, err
	}

	data := make([]byte, packetLength)

	if _, err := r.Read(data); err != nil {
		return 0, nil, err
	}

	buf := bytes.NewBuffer(data)

	packetType, err := readVarInt(buf)

	if err != nil {
		return 0, nil, err
	}

	return packetType, buf, nil
}

func writePacket(w io.Writer, packetType int32, buf *bytes.Buffer) error {
	data := &bytes.Buffer{}

	if err := writeVarInt(data, packetType); err != nil {
		return err
	}

	if _, err := io.Copy(data, buf); err != nil {
		return err
	}

	if err := writeVarInt(w, int32(data.Len())); err != nil {
		return err
	}

	if _, err := io.Copy(w, data); err != nil {
		return err
	}

	return nil
}
