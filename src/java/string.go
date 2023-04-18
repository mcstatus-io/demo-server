package java

import "io"

func readString(r io.Reader) (string, error) {
	length, err := readVarInt(r)

	if err != nil {
		return "", err
	}

	data := make([]byte, length)

	if _, err := r.Read(data); err != nil {
		return "", err
	}

	return string(data), nil
}

func writeString(w io.Writer, value string) (err error) {
	if err = writeVarInt(w, int32(len(value))); err != nil {
		return err
	}

	_, err = w.Write([]byte(value))

	return
}
