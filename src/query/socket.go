package query

import (
	"bytes"
	"fmt"
	"net"
	"time"

	"main/src/config"
)

var (
	socket   net.PacketConn   = nil
	conf     *config.Config   = nil
	sessions map[int32]string = make(map[int32]string)
)

// Listen creates a new TCP socket server using the address specified in the configuration file.
func Listen(c *config.Config) (err error) {
	conf = c

	socket, err = net.ListenPacket("udp", fmt.Sprintf("%s:%d", c.JavaEdition.Query.Host, c.JavaEdition.Query.Port))

	if err == nil {
		go func() {
			for {
				for k := range sessions {
					delete(sessions, k)
				}

				time.Sleep(conf.JavaEdition.Query.GlobalSessionExpiration)
			}
		}()
	}

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

	packetType, sessionID, err := readBasePacket(r)

	if err != nil {
		return
	}

	buf := &bytes.Buffer{}

	switch packetType {
	case 0x09: // Generate challenge token
		{
			if err = writeHandshakePacket(buf, sessionID); err != nil {
				return
			}

			break
		}
	case 0x00: // Request
		{
			isFullStat, err := readRequestPacket(r, buf, sessionID)

			if err != nil {
				return
			}

			if isFullStat {
				if err = writeFullStatPacket(buf, sessionID); err != nil {
					return
				}
			} else {
				if err = writeBasicStatPacket(buf, sessionID); err != nil {
					return
				}
			}

			break
		}
	}

	socket.WriteTo(buf.Bytes(), addr)
}
