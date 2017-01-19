package models

import (
	"strconv"
	"time"

	redis "gopkg.in/redis.v5"

	database "github.com/NorbertKa/GoToken/databases"
)

type Token string

type Tokens []Token

func (t Token) Create(db *database.Redis) error {
	timeNow := time.Now().UnixNano()
	z := redis.Z{
		Score:  float64(timeNow),
		Member: string(t),
	}
	err := db.ZAdd("IssuedTokens", z).Err()
	if err != nil {
		return err
	}
	err = db.Set(strconv.Itoa(int(timeNow))+"-token", t, 0).Err()
	return nil
}
