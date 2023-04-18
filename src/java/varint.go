package java

import (
	"errors"
	"io"
)

func readVarInt(r io.Reader) (int32, error) {
	var result int32 = 0
	var numRead int = 0

	for {
		data := make([]byte, 1)

		n, err := r.Read(data)

		if err != nil {
			return 0, err
		}

		if n < 1 {
			return 0, io.EOF
		}

		value := (data[0] & 0b01111111)
		result |= int32(value) << (7 * numRead)

		numRead++

		if numRead > 5 {
			return 0, errors.New("varint too big")
		}

		if (data[0] & 0b10000000) == 0 {
			break
		}
	}

	return result, nil
}

func writeVarInt(w io.Writer, val int32) error {
	for {
		if (uint32(val) & 0xFFFFFF80) == 0 {
			_, err := w.Write([]byte{byte(val)})

			return err
		}

		if _, err := w.Write([]byte{byte(val&0x7F | 0x80)}); err != nil {
			return err
		}

		val = int32(uint32(val) >> 7)
	}
}
