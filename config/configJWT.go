package config

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
)

//nambain tokenConfig untuk login
type TokenConfig struct {
	ApplicationName  string
	JwtSigningMethod *jwt.SigningMethodHMAC
	JwtSignatureKey string 
	AccessTokenLifeTime time.Duration
	Client *redis.Client
}


type ConfigJWT struct {
	TokenConfig //tambain jwt token config 

}

func (c ConfigJWT) readConfigJWT() ConfigJWT {
	c.TokenConfig = TokenConfig{
	ApplicationName:  "Enigma",
	JwtSigningMethod: jwt.SigningMethodHS256,
	JwtSignatureKey: "31N!GMA",
	AccessTokenLifeTime: 60 * time.Second,
	Client: redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}),
		
	}
	return c
}

func NewConfigJWT() ConfigJWT {
	cfg := ConfigJWT{}
	return cfg.readConfigJWT()
}

