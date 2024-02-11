package vote

import (
	"crypto/rand"
	"encoding/hex"
)

func generateChallenge() string {
	data := make([]byte, 8)

	if _, err := rand.Read(data); err != nil {
		panic(err)
	}

	return hex.EncodeToString(data)
}
