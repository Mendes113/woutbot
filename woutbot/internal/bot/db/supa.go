package db

import (
	"context"
	"database/sql"
	"fmt"
	
	"woutbot/config"

	// Importe o driver SQL do Supabase.
	_ "github.com/lib/pq"
	"github.com/supabase-community/supabase-go"
)

func Connect() (*supabase.Client, error) {
	API_URL := config.GetURLDB()
	API_KEY := config.GetAPIDB()

	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize client: %w", err)
	}


	return client, nil
}

func SaveToDB() error {
	API_URL := config.GetURLDB()
	API_KEY := config.GetAPIDB()

	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		return fmt.Errorf("cannot initialize client: %w", err)
	}

	data := []map[string]interface{}{
		{"name": "Brazil", "code": "BR"},
		{"name": "United States", "code": "US"},
		{"name": "India", "code": "IN"},
	}
	_, _, err = client.From("countries").Insert(data, true, "exact", "false", "").Execute()
	if err != nil {
		return fmt.Errorf("error executing insert query: %w", err)
	}

	fmt.Println("Data inserted successfully")

	return nil
}



func CreateTable(db *sql.DB) error {
	// Consulte a documentação do Supabase para determinar a sintaxe correta da consulta SQL para criar sua tabela.
	// Aqui está um exemplo simples de criação de tabela:
	query := `
		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE
		);
	`

	// Execute a consulta SQL usando o método `Exec()` da conexão com o banco de dados.
	_, err := db.ExecContext(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}