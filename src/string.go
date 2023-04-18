package main

import "io"

func ReadString(r io.Reader) (string, error) {
	length, _, err := ReadVarInt(r)

	if err != nil {
		return "", err
	}

	data := make([]byte, length)

	if _, err := r.Read(data); err != nil {
		return "", err
	}

	return string(data), nil
}

func WriteString(value string, w io.Writer) error {
	if _, err := WriteVarInt(int32(len(value)), w); err != nil {
		return err
	}

	_, err := w.Write([]byte(value))

	return err
}
