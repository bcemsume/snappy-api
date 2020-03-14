package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	DBHOST         = "DBHOST"
	DBPORT         = "DBPORT"
	DBUSER         = "DBUSER"
	DBPASS         = "DBPASS"
	DBNAME         = "DBNAME"
	SSLMODE        = "SSLMODE"
	LOG_FILE_PATH  = "LOG_FILE_PATH"
	LOG_FILE_NAME  = "LOG_FILE_NAME"
	TOKEN_PASSWORD = "TOKEN_PASSWORD"
	LOG_FILE_SIZE  = "LOG_FILE_SIZE"
)

func DBConfigs() map[string]string {
	conf := make(map[string]string)
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
		return conf

	}
	host := os.Getenv(DBHOST)
	port := os.Getenv(DBPORT)
	user := os.Getenv(DBUSER)
	password := os.Getenv(DBPASS)
	name := os.Getenv(DBNAME)
	sslMode := os.Getenv(SSLMODE)

	conf[DBHOST] = host
	conf[DBPORT] = port
	conf[DBUSER] = user
	conf[DBPASS] = password
	conf[DBNAME] = name
	conf[SSLMODE] = sslMode
	return conf
}

func LogConfigs() map[string]string {
	conf := make(map[string]string)
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
		return conf
	}
	conf[LOG_FILE_PATH] = os.Getenv(LOG_FILE_PATH)
	conf[LOG_FILE_NAME] = os.Getenv(LOG_FILE_NAME)
	conf[LOG_FILE_SIZE] = os.Getenv(LOG_FILE_SIZE)
	return conf
}

func GetTokenConfig() map[string]string {

	conf := make(map[string]string)

	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
		return conf
	}

	conf[TOKEN_PASSWORD] = os.Getenv(TOKEN_PASSWORD)

	return conf
}
