package vote

import (
	"fmt"
	"net"

	"main/src/config"
)

var (
	socket net.Listener   = nil
	conf   *config.Config = nil
)

// Listen creates a new TCP socket server using the address specified in the configuration file.
func Listen(c *config.Config) (err error) {
	conf = c

	socket, err = net.Listen("tcp", fmt.Sprintf("%s:%d", c.Votifier.Host, c.Votifier.Port))

	return
}

// Close closes the socket server.
func Close() error {
	return socket.Close()
}

// AcceptConnections should be started in a Goroutine and accepts new connections from the socket server.
func AcceptConnections() {
	for {
		conn, err := socket.Accept()

		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	challenge := generateChallenge()

	if err := sendHandshake(conn, challenge); err != nil {
		writeError(conn, err.Error())

		return
	}

	if err := readPayload(conn, challenge); err != nil {
		writeError(conn, err.Error())

		return
	}

	if err := writeResponse(conn); err != nil {
		writeError(conn, err.Error())

		return
	}
}
