package models

import (
	"errors"

	"github.com/NorbertKa/GoToken/config"
	"github.com/NorbertKa/GoToken/databases"
	"golang.org/x/crypto/bcrypt"
)

type Userprofile struct {
	Id          int    `json:"id"`
	Created     string `json:"created"`
	LastChanges string `json:"lastChanges"`
	LastLogin   string `json:"lastLogin"`
	Identifier  string `json:"identifier"`
	Password    string `json:"password"`
}

const (
	ErrFailedToGenerateHash string = "Error generating hash from password"
)

func (u Userprofile) Create(db *database.Postgre, conf *config.Config) error {
	stmt, err := db.Prepare("INSERT INTO gotoken.userprofile(identifier, password)")
	if err != nil {
		return err
	}
	hash := []byte{}
	hash, err = bcrypt.GenerateFromPassword([]byte(u.Password), conf.HashCost)
	if err != nil {
		return errors.New(ErrFailedToGenerateHash)
	}
	_, err = stmt.Exec(u.Identifier, string(hash))
	if err != nil {
		return err
	}
	return nil
}
