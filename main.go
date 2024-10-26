package main

import (
	"github.com/ChampionBuffalo1/vessel/cli"
	"github.com/ChampionBuffalo1/vessel/pkg/log"
)

func main() {
	log.InitLogger()
	cli.Execute()
}
