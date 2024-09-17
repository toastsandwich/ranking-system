package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	SERVERHOST       string
	SERVERPORT       string
	DATABASEHOST     string
	DATABASEPORT     string
	DATABASEDB       string
	DATABASEUSERNAME string
	DATABASEPASSWORD string
	LOGPATH          string
	DB               *sql.DB
)

func ConfigureDB() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   DATABASEUSERNAME,
		Passwd: DATABASEPASSWORD,
		Net:    "tcp",
		Addr:   DATABASEHOST + ":" + DATABASEPORT,
		DBName: DATABASEDB,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, err
}

func ReadConfigFile() error {
	viper.SetConfigName("config.toml")
	viper.AddConfigPath("resources/")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	SERVERHOST = viper.GetString("server.host")
	SERVERPORT = viper.GetString("server.port")
	DATABASEDB = viper.GetString("database.db")
	DATABASEHOST = viper.GetString("database.host")
	DATABASEPORT = viper.GetString("database.port")
	DATABASEPASSWORD = viper.GetString("database.password")
	DATABASEUSERNAME = viper.GetString("database.username")
	LOGPATH = viper.GetString("log.path")
	return nil
}

func CreateLogFile() (*os.File, error) {
	file, err := os.OpenFile(LOGPATH, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return file, nil
}

func init() {
	var err error
	err = ReadConfigFile()
	if err != nil {
		log.Fatal(err)
	}
	db, err := ConfigureDB()
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	query := `CREATE TABLE IF NOT EXISTS scores (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(100) NOT NULL UNIQUE,
    country VARCHAR(50) NOT NULL DEFAULT 'anonymous',
    state VARCHAR(50) NOT NULL DEFAULT 'anonymous',
    score FLOAT NOT NULL
    );`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
