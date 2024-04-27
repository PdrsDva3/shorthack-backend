package config

import (
	"fmt"
	"os"
	"path/filepath"
)

import (
	"github.com/spf13/viper"
)

const (
	DBUser     = "DB_USER"
	DBPassword = "DB_PASSWORD"
	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
	DBName     = "DB_NAME"
	TimeOut    = "TIMEOUT"
)

func InitConfig() {
	path, _ := os.Getwd()

	path = filepath.Join(path, "..")
	//path = filepath.Join(path, "shorthack_backend")

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	fmt.Println(viper.GetString(DBHost))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Config initialization error: %v", err.Error()))
	}
}
