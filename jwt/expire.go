package jwt

import (
	"errors"
	"time"
)

const (
	ErrTokenExpired string = "JWT Token is expired"
)

func (t TokenInfo) CheckIfExpired() (bool, error) {
	timeNow := time.Now().UnixNano()
	if (t.Nbf + t.Exp) > int(timeNow) {
		return true, errors.New(ErrTokenExpired)
	} else {
		return false, nil
	}
}