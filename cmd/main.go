package main

import (
	"log"

	"github.com/teamcubation/teamcubot-message/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
