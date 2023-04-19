package java

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"os"

	"main/src/config"
	"main/src/util"
)

var (
	socket  net.Listener   = nil
	conf    *config.Config = nil
	favicon string         = ""
)

func init() {
	if stat, err := os.Stat("server-icon.png"); err == nil && !stat.IsDir() {
		data, err := os.ReadFile("server-icon.png")

		if err != nil {
			log.Fatal(err)
		}

		favicon = "data:image/png;base64," + base64.RawStdEncoding.EncodeToString(data)
	}
}

// Listen creates a new TCP socket server using the address specified in the configuration file.
func Listen(c *config.Config) (err error) {
	conf = c

	socket, err = net.Listen("tcp", fmt.Sprintf("%s:%d", c.JavaEdition.Status.Host, c.JavaEdition.Status.Port))

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

	nextState, err := readHandshakePacket(conn)

	if err != nil {
		return
	}

	switch nextState {
	case nextStateStatus:
		{
			var payload int64

			if err = readRequestPacket(conn); err != nil {
				return
			}

			if err = writeResponsePacket(conn); err != nil {
				return
			}

			if payload, err = readPingPacket(conn); err != nil {
				return
			}

			if err = writePongPacket(conn, payload); err != nil {
				return
			}

			return
		}
	case nextStateLogin:
		{
			uuid, err := readLoginStartPacket(conn)

			if err != nil {
				return
			}

			if err = writeDisconnectPacket(conn, conf.JavaEdition.Status.DisconnectReason); err != nil {
				return
			}

			if uuid != nil {
				go util.AddSamplePlayer(*uuid)
			}

			return
		}
	}
}
