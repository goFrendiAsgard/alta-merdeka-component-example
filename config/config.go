package config

import "os"

type Cfg struct {
	Port               string
	DBConnectionString string
}

func NewCfg() Cfg {
	// port
	port := os.Getenv("APP_HTTP_PORT")
	if port == "" {
		port = "3000"
	}
	// connection string
	dbConnectionString := os.Getenv("APP_DB_CONNECTION_STRING")
	if dbConnectionString == "" {
		dbConnectionString = "root:toor@tcp(localhost:3306)/alta?charset=utf8mb4&parseTime=True&loc=Local"
	}
	// config
	return Cfg{
		Port:               port,
		DBConnectionString: dbConnectionString,
	}
}
