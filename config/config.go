package config

import (
	"fmt"
	"os"
)

type ApiConfig struct {
	Url string
}

type DbConfig struct {
	DataSourceName string
}

type Config struct {
	ApiConfig
	DbConfig

	
}

// //nambain tokenConfig untuk login
// type TokenConfig struct {
// 	ApplicationName  string
// 	JwtSigningMethod *jwt.SigningMethodHMAC
// 	JwtSignatureKey string 
// 	AccessTokenLifeTime time.Duration
// }

func (c *Config) readConfig()  {
	api := os.Getenv("API_URL")
	//set API_URL=localhost:8888
	// c.ApiConfig = ApiConfig{Url: api}

	dbHost := os.Getenv("DB_HOST")         //set DB_HOST=localhost
	dbPort := os.Getenv("DB_PORT")         //set DB_PORT=5432
	dbUser := os.Getenv("DB_USER")         //set DB_USER=postgres
	dbPassword := os.Getenv("DB_PASSWORD") //set DB_PASSWORD=12345678
	dbName := os.Getenv("DB_NAME")         //set DB_NAME=db_gin_latihan

	// urutan url koneksi ke db postgres buat gorm
	// localhost:postgres@12345678:db_enigma_shop_v2/5432

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	c.DbConfig = DbConfig{DataSourceName: dsn}
	c.ApiConfig = ApiConfig{Url: api}

	//nambain token config
	// c.TokenConfig = TokenConfig{
	// 	ApplicationName:  "Enigma",
	// 	JwtSigningMethod: jwt.SigningMethodHS256,
	// 	JwtSignatureKey: "31N!GMA",
	// 	AccessTokenLifeTime: 60 * time.Second,
			
	// 	}
	// return c
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}