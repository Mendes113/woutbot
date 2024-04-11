package config

// Coloque aqui as estruturas e funções relacionadas à configuração do bot


import (
	
	"log"
	"os"
)


func GetAPI() string {

	api := os.Getenv("API")
	if api == "" {
		log.Fatal("Chave da API não encontrada ")
	}

	return api
}

