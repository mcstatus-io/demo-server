package bedrock

import (
	"bytes"
	"fmt"
	"net"

	"main/src/config"
)

var (
	socket net.PacketConn = nil
	conf   *config.Config = nil
)

// Listen creates a new TCP socket server using the address specified in the configuration file.
func Listen(c *config.Config) (err error) {
	conf = c

	socket, err = net.ListenPacket("udp", fmt.Sprintf("%s:%d", c.BedrockEdition.Status.Host, c.BedrockEdition.Status.Port))

	return
}

// Close closes the socket server.
func Close() error {
	return socket.Close()
}

// AcceptConnections should be started in a Goroutine and accepts new connections from the socket server.
func AcceptConnections() {
	for {
		data := make([]byte, 4096)

		n, addr, err := socket.ReadFrom(data)

		if err != nil {
			continue
		}

		go handlePacket(data[:n], addr)
	}
}

func handlePacket(data []byte, addr net.Addr) {
	r := bytes.NewReader(data)

	if err := readUnconnectedPingPacket(r); err != nil {
		return
	}

	buf := &bytes.Buffer{}

	if err := writeUnconnectedPongPacket(buf); err != nil {
		return
	}

	socket.WriteTo(buf.Bytes(), addr)
}
