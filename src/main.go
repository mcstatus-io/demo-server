package main

import (
	"log"
	"os"
	"os/signal"

	"main/src/config"
	"main/src/java"
	"main/src/query"
	"main/src/util"
)

var (
	conf *config.Config = &config.Config{}
)

func init() {
	var err error

	if err = conf.ReadFile("config.yml"); err != nil {
		log.Fatal(err)
	}

	if conf.JavaStatus.Enable {
		if err = java.Listen(conf); err != nil {
			log.Fatal(err)
		}

		log.Printf("Listening for Java Edition statuses on %s:%d\n", conf.JavaStatus.Host, conf.JavaStatus.Port)
	}

	if conf.Query.Enable {
		if err = query.Listen(conf); err != nil {
			log.Fatal(err)
		}

		log.Printf("Listening for query connections on %s:%d\n", conf.Query.Host, conf.Query.Port)
	}

	util.SetSamplePlayers(conf.Players.Sample)
}

func main() {
	if conf.JavaStatus.Enable {
		defer java.Close()

		go java.AcceptConnections()
	}

	if conf.Query.Enable {
		defer query.Close()

		go query.AcceptConnections()
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	<-s
}
