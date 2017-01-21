package models

import (
	"github.com/NorbertKa/GoToken/config"
	database "github.com/NorbertKa/GoToken/databases"
	"github.com/NorbertKa/GoToken/jwt"
	"gopkg.in/redis.v5"
)

func BanToken(db *database.Redis, conf *config.Config, token string) error {
	t, err := jwt.DecodeToken(token, conf)
	if err != nil {
		return err
	}

	isExpired, err := t.CheckIfExpired()
	if isExpired {
		return err
	}

	z := redis.Z{
		Score:  float64(t.Nbf),
		Member: token,
	}

	err = db.ZAdd("bannedTokens", z).Err()
	if err != nil {
		return err
	}
	return nil
}
