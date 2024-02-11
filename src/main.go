package main

import (
	"log"
	"os"
	"os/signal"

	"main/src/bedrock"
	"main/src/config"
	"main/src/java"
	"main/src/query"
	"main/src/util"
	"main/src/vote"
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

	if conf.BedrockEdition.Status.Enable {
		if err = bedrock.Listen(conf); err != nil {
			log.Fatal(err)
		}

		log.Printf("Listening for Bedrock Edition statuses on %s:%d\n", conf.BedrockEdition.Status.Host, conf.BedrockEdition.Status.Port)
	}

	if conf.JavaEdition.Query.Enable {
		if err = query.Listen(conf); err != nil {
			log.Fatal(err)
		}

		log.Printf("Listening for query connections on %s:%d\n", conf.JavaEdition.Query.Host, conf.JavaEdition.Query.Port)
	}

	if conf.Votifier.Enable {
		if err = vote.Listen(conf); err != nil {
			log.Fatal(err)
		}

		log.Printf("Listening for Votifier on %s:%d\n", conf.Votifier.Host, conf.Votifier.Port)
	}

	util.SetSamplePlayers(conf.JavaEdition.Options.Players.Sample)
}

func main() {
	if conf.JavaEdition.Status.Enable {
		defer java.Close()

		go java.AcceptConnections()
	}

	if conf.BedrockEdition.Status.Enable {
		defer bedrock.Close()

		go bedrock.AcceptConnections()
	}

	if conf.JavaEdition.Query.Enable {
		defer query.Close()

		go query.AcceptConnections()
	}

	if conf.Votifier.Enable {
		defer vote.Close()

		go vote.AcceptConnections()
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	<-s
}
