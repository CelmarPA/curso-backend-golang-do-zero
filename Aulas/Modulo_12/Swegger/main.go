package main

import (
	"github.com/CelmarPA/curso-backend-golang-do-zero/Aulas/Modulo_12/Swegger/api"
	"github.com/rs/zerolog/log"
)

func main() {
	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")

	}
}
