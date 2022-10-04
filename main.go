package main

import (
	config "api-jogos-twitch/config/database"
	"api-jogos-twitch/interfaces/swagger"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	var (
		configuration config.Configurations
		database      *sql.DB
		erro          error
	)

	viper.SetConfigName("dev")

	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.SetConfigType("env")

	if erro = viper.ReadInConfig(); erro != nil {
		zap.L().Error("Erro ao ler as configurações", zap.Error(erro))
		return
	}

	if erro = viper.Unmarshal(&configuration); erro != nil {
		zap.L().Error("Erro ao decodificar as configurações", zap.Error(erro))
		return
	}

	if database, erro = config.AbrirConexao(configuration); erro != nil {
		zap.L().Error("Erro ao abrir conexão com o banco de dados", zap.Error(erro))
		return
	}
	defer database.Close()

	router := gin.New()

	r := router.Group("/")
	swagger.Router(r)

	zap.L().Info(os.Getenv("PORT"))
	zap.L().Info(configuration.Port)

	log.Fatal(http.ListenAndServe(configuration.Port, router))
}
