package main

import (
	"github.com/brunetto/goutils/system"
	"log"
)

func main() {
	var err error

	err = system.MonitorAndKill("parentalcontrol")

	if err != nil {
		log.Fatal(err)
	}
}
