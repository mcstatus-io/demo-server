package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	config         *Config      = &Config{}
	statusListener net.Listener = nil
	queryListener  net.Listener = nil
	serverIcon     *string      = nil
)

func init() {
	var err error

	if err = config.ReadFile("config.yml"); err != nil {
		log.Fatal(err)
	}

	if !config.JavaStatus.Enable && !config.Query.Enable {
		log.Fatal("either status.enable or query.enable must be true in config.yml")
	}

	if config.JavaStatus.Enable {
		if statusListener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", config.JavaStatus.Host, config.JavaStatus.Port)); err != nil {
			log.Fatal(err)
		}

		log.Printf("Now listening on %s:%d for status\n", config.JavaStatus.Host, config.JavaStatus.Port)
	}

	if config.Query.Enable {
		if queryListener, err = net.Listen("udp", fmt.Sprintf("%s:%d", config.Query.Host, config.Query.Port)); err != nil {
			log.Fatal(err)
		}

		log.Printf("Now listening on %s:%d for query\n", config.Query.Host, config.Query.Port)
	}

	if stat, err := os.Stat("server-icon.png"); err == nil && !stat.IsDir() {
		data, err := os.ReadFile("server-icon.png")

		if err != nil {
			log.Fatal(err)
		}

		value := "data:image/png;base64," + base64.RawStdEncoding.EncodeToString(data)

		serverIcon = &value
	}
}

func main() {
	if config.JavaStatus.Enable {
		go AcceptJavaStatusConnections()
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	<-s
}
