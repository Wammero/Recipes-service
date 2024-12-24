package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func LoadDatabaseConnectionString() (string, error) {
	viper.SetConfigFile("/app/config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
		return "", err
	}

	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	dbname := viper.GetString("database.dbname")
	sslmode := viper.GetString("database.sslmode")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbname, sslmode)

	return connStr, nil
}
