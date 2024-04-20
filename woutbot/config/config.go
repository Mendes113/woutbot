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


func GetURLDB() string {


	url := os.Getenv("URLDB")
	if url == "" {
		log.Fatal("URL do banco de dados não encontrada ")
	}

	return url
}


func GetAPIDB() string {

	api := os.Getenv("APIDB")
	if api == "" {
		log.Fatal("Chave da API do banco de dados não encontrada ")
	}

	return api
}
