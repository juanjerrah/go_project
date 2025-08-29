package main

import (
	"log"

	"github.com/juanjerrah/go-project/internal/config"
	"github.com/juanjerrah/go-project/internal/server"
)

func main() {
	cfg := config.LoadConfig()

	server := server.NewServer(cfg)

	if err := server.InitializeDatabase(); err != nil {
		panic(err)
	}

	if err := server.Run(); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)	}
}