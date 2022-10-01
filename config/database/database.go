package config

import (
	"api-jogos-twitch/oops"
	"database/sql"

	_ "github.com/lib/pq"
)

// AbrirConexao abre conexão com o banco de dados
func AbrirConexao(configuration Configurations) (database *sql.DB, erro error) {
	if database, erro = sql.Open("postgres",
		"host="+configuration.Database.DBHost+
			" port="+configuration.Database.DBPort+
			" user="+configuration.Database.DBUser+
			" password="+configuration.Database.DBPassword+
			" dbname="+configuration.Database.DBName+
			" sslmode=disable"); erro != nil {
		return nil, oops.Err(erro)
	}

	return
}
