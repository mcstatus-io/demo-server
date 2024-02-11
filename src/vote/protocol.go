package vote

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
)

type voteMessage struct {
	Payload   string `json:"payload"`
	Signature string `json:"signature"`
}

type votePayload struct {
	ServiceName string `json:"serviceName"`
	Username    string `json:"username"`
	Timestamp   int64  `json:"timestamp"`
	Challenge   string `json:"challenge"`
	UUID        string `json:"uuid"`
}

type voteResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func sendHandshake(w io.Writer, challenge string) error {
	_, err := w.Write([]byte(fmt.Sprintf("VOTIFIER 2 %s\n", challenge)))

	return err
}

func readPayload(r io.Reader, challenge string) error {
	var (
		identifier    int16
		messageLength int16
	)

	if err := binary.Read(r, binary.BigEndian, &identifier); err != nil {
		return err
	}

	if identifier != 0x733A {
		return fmt.Errorf("invalid identifier: %X", identifier)
	}

	if err := binary.Read(r, binary.BigEndian, &messageLength); err != nil {
		return err
	}

	data := make([]byte, messageLength)

	if _, err := r.Read(data); err != nil {
		return err
	}

	var message voteMessage

	if err := json.Unmarshal(data, &message); err != nil {
		return err
	}

	var payload votePayload

	if err := json.Unmarshal([]byte(message.Payload), &payload); err != nil {
		return err
	}

	if payload.Challenge != challenge {
		return fmt.Errorf("unexpected challenge: %s", payload.Challenge)
	}

	hash := hmac.New(sha256.New, []byte(conf.Votifier.Token))
	hash.Write([]byte(message.Payload))

	if base64.StdEncoding.EncodeToString(hash.Sum(nil)) != message.Signature {
		return fmt.Errorf("invalid message signature (is the token wrong?): %s", message.Signature)
	}

	return nil
}

func writeResponse(w io.Writer) error {
	data, err := json.Marshal(voteResponse{
		Status: "ok",
	})

	if err != nil {
		return err
	}

	_, err = w.Write(append(data, '\n'))

	return err
}

func writeError(w io.Writer, message string) error {
	data, err := json.Marshal(voteResponse{
		Status: "error",
		Error:  message,
	})

	if err != nil {
		return err
	}

	_, err = w.Write(append(data, '\n'))

	return err
}
