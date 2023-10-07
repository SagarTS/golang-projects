package infrastructure

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	PORT,_ := strconv.Atoi(os.Getenv("DB_PORT"))

	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     PORT,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
	}
	return &dbConfig
}

// databse configuratuon
func DbURL(dbConfig *DBConfig) string {
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	return url
}
