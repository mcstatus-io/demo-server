package query

import "io"

func writeNTString(w io.Writer, value string) error {
	if _, err := w.Write([]byte(value)); err != nil {
		return err
	}

	_, err := w.Write([]byte{0x00})

	return err
}
