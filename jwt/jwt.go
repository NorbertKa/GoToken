package jwt

import (
	"time"

	"github.com/NorbertKa/GoToken/config"
	"github.com/NorbertKa/GoToken/models"
	"github.com/dgrijalva/jwt-go"
)

type TokenInfo struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
	Nbf    int `json:"nbf"`
	Exp    int `json:"exp"`
}

func DecodeToken(token string, conf *config.Config) (*TokenInfo, error) {
	claims := TokenInfo{}
	tok, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return conf.Secret, nil
	})
	if tok.Signature != "HS512" {
		return nil, jwt.ErrSignatureInvalid
	}
	if err != nil {
		return nil, err
	}
	return &claims, nil
}

func EncodeToken(userProfile *models.Userprofile, conf *config.Config) (string, error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"userId":     userProfile.Id,
		"identifier": userProfile.Identifier,
		"nbf":        time.Now().UnixNano(),
		"exp":        time.Minute.Nanoseconds() * 15,
	})
	tokenString, err := tok.SignedString(conf.Secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
