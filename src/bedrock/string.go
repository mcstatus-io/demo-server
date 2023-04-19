package bedrock

import (
	"encoding/binary"
	"io"
)

func writeString(w io.Writer, value string) error {
	if err := binary.Write(w, binary.BigEndian, uint16(len(value))); err != nil {
		return err
	}

	_, err := w.Write([]byte(value))

	return err
}
