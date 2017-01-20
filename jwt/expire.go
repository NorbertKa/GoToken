package jwt

import (
	"errors"
	"time"
)

const (
	ErrTokenExpired string = "JWT Token is expired"
)

func (t TokenInfo) CheckIfExpired() (bool, error) {
	if err != nil {
		return false, err
	}
	timeNow := time.Now().UnixNano()
	if (t.Nbf + t.Exp) > int(timeNow) {
		return true, errors.New(ErrTokenExpired)
	} else {
		return false, nil
	}
}
