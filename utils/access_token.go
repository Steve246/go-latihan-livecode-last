package utils

import (
	"fmt"
	"go_livecode_persiapan/config"
	"go_livecode_persiapan/model"
	"time"

	"github.com/golang-jwt/jwt"
)

type Token interface {
	CreateAccessToken(cred *model.Credential) (string, error)

	VerifyAccessToken(tokenString string) (jwt.MapClaims, error)

	// StoreAccessToken(userName string, tokenDetail *model.TokenDetails) error

	// FetchAccessToken(accessDetail *model.AccessDetail) (string, error)
}

type token struct {
	cfg config.TokenConfig
}

// func (t *token) StoreAccessToken(userName string, tokenDetail *model.TokenDetails) error {
// 	at := time.Unix(tokenDetail.AtExpires, 0)
// 	now := time.Now()
// 	err := t.cfg.Client.Set(context.Background(), tokenDetail.AccessUuid, userName, at.Sub(now)).Err()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (t *token) FetchAccessToken(accessDetail *model.AccessDetail) (string, error) {
// 	if accessDetail != nil {
// 		result, err := t.cfg.Client.Get(context.Background(), accessDetail.AccessUuid).Result()
// 		if err != nil {
// 			return "", err
// 		}
// 		return result, nil
// 	} else {
// 		return "", errors.New("invalid access")
// 	}
// }

func (t *token) CreateAccessToken(cred *model.Credential) (string, error) {
	now := time.Now().UTC()
	end := now.Add(t.cfg.AccessTokenLifeTime)

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: t.cfg.ApplicationName,
		},
		Username: cred.Username,
		Email:    cred.Email,
	}

	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = end.Unix()

	token := jwt.NewWithClaims(
		t.cfg.JwtSigningMethod,
		claims,
	)

	return token.SignedString([]byte(t.cfg.JwtSignatureKey))

}

func (t *token) VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != t.cfg.JwtSigningMethod {
			return nil, fmt.Errorf("signing mehod invalid")
		}
		return []byte(t.cfg.JwtSignatureKey), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || token.Valid || claims["iss"] != t.cfg.ApplicationName {
		return nil, err
	}
	return claims, nil
}

func NewTokenService(cfg config.TokenConfig) Token {
	return &token{cfg: cfg}
}

