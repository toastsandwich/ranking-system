package main

import (
	"log"

	"github.com/toastsandwich/realtime-ranking-system/config"
	"github.com/toastsandwich/realtime-ranking-system/server"
)

func main() {
	file, err := config.CreateLogFile()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	addr := config.SERVERHOST + ":" + config.SERVERPORT
	server := server.NewApiServer(addr)
	log.Fatal(server.Start())
}
