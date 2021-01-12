package main

import (
	"clly/apterture/internal/daemon"
	"log"
)

func main() {
	daemon := daemon.Aperture{}
	if err := daemon.Run(); err != nil {
		log.Fatal("fatal error", err)
	}
}
