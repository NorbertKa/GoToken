package jwt

import (
	"time"
	"errors"

	"github.com/NorbertKa/GoToken/config"
	"github.com/NorbertKa/GoToken/models"
	"github.com/dgrijalva/jwt-go"
)

const (
	ErrUndefinedAlgo string = "Undefined encoding algorithm used"
	ErrUnexpectedSigningMethod string = "Unexpected signing method"
	ErrInvalidToken string = "Invalid token"
)

type Claims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
}

func DecodeToken(token string, conf *config.Config) (*Claims, error) {
	claims := Claims{}
	switch conf.Algorithm {
	case "HS256", "HS384", "HD512":
		t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New(ErrUnexpectedSigningMethod)
			}
			return []byte(conf.Secret), nil
		})
		if err != nil {
			return nil, err
		}
		if !t.Valid {
			return nil, errors.New(ErrInvalidToken)
		}
		return &claims, nil
	/*
	TODO: Implement the rest of encryption algorithms.
	case "ES256", "ES384", "ES512":
	case "RS256", "RS384", "RS384":
	case "PS256", "PS384", "PS512":
	*/
	default:
		return nil, errors.New(ErrUndefinedAlgo)
	}
}


func hsEncode(claims *Claims, algo string, conf *config.Config) (*jwt.Token, error) {
	var token *jwt.Token
	switch algo {
	case "HS256":
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token, nil
	case "HS384":
		token = jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
		return token, nil
	case "HS512":
		token = jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		return token, nil
	default:
		return nil, errors.New(ErrUndefinedAlgo)
	}
}

func EncodeToken(userProfile *models.Userprofile, conf *config.Config) (string, error) {
	timeNow := time.Now()
	timeNowNano := timeNow.UnixNano()
	claims := Claims{
		jwt.StandardClaims{
			Issuer:    conf.Issuer,
			IssuedAt: timeNowNano,
			NotBefore: timeNowNano,
			ExpiresAt: timeNow.Add(time.Minute * time.Duration(conf.TokenDuration)).UnixNano(),
			Id: userProfile.Identifier,
		},
		userProfile.Id,
	}
	var signedString string
	switch conf.Algorithm {
	case "HS256", "HS384", "HD512":
		token, err := hsEncode(&claims, conf.Algorithm, conf)
		if err != nil {
			return nil, err
		}
		signedString, err = token.SignedString(conf.Secret)
		if err != nil {
			return nil, err
		}
	/*
	TODO: Implement the rest of encryption algorithms.
	case "RS256", "RS384", "RS384":
	case "PS256", "PS384", "PS512":
	*/
	default:
	return nil, errors.New(ErrUndefinedAlgo)
	}
	return signedString, nil
}
