package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// TODO config reading change to configor package
// DBHOST s
const (
	DBHOST        = "DBHOST"
	DBPORT        = "DBPORT"
	DBUSER        = "DBUSER"
	DBPASS        = "DBPASS"
	DBNAME        = "DBNAME"
	SSLMODE       = "SSLMODE"
	LOGFILEPATH   = "LOG_FILE_PATH"
	LOGFILENAME   = "LOG_FILE_NAME"
	TOKENPASSWORD = "TOKEN_PASSWORD"
	LOGFILESIZE   = "LOG_FILE_SIZE"
	DEBUGMODE     = "DEBUG_MODE"
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
	conf[LOGFILEPATH] = os.Getenv(LOGFILEPATH)
	conf[LOGFILENAME] = os.Getenv(LOGFILENAME)
	conf[LOGFILESIZE] = os.Getenv(LOGFILESIZE)
	return conf
}

func GetTokenConfig() map[string]string {

	conf := make(map[string]string)

	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
		return conf
	}

	conf[TOKENPASSWORD] = os.Getenv(TOKENPASSWORD)

	return conf
}

func GetAppConfig() map[string]string {

	conf := make(map[string]string)

	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
		return conf
	}

	conf[DEBUGMODE] = os.Getenv(DEBUGMODE)

	return conf
}
