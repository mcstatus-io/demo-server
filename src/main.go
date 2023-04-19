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

	if conf.JavaEdition.Status.Enable {
		if err = java.Listen(conf); err != nil {
			log.Fatal(err)
		}

		log.Printf("Listening for Java Edition statuses on %s:%d\n", conf.JavaEdition.Status.Host, conf.JavaEdition.Status.Port)
	}

	if conf.JavaEdition.Query.Enable {
		if err = query.Listen(conf); err != nil {
			log.Fatal(err)
		}

		log.Printf("Listening for query connections on %s:%d\n", conf.JavaEdition.Query.Host, conf.JavaEdition.Query.Port)
	}

	util.SetSamplePlayers(conf.JavaEdition.Options.Players.Sample)
}

func main() {
	if conf.JavaEdition.Status.Enable {
		defer java.Close()

		go java.AcceptConnections()
	}

	if conf.JavaEdition.Query.Enable {
		defer query.Close()

		go query.AcceptConnections()
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	<-s
}
