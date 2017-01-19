package config

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	multierror "github.com/hashicorp/go-multierror"
	_ "github.com/lib/pq"
)

const (
	ErrInvalidTimeout      string = "Invalid DB timeout"
	ErrInvalidSslMode      string = "Invalid DB SSL mode"
	ErrCantOpenSslCert     string = "Cant open DB SSL Cert file"
	ErrCantOpenSslKey      string = "Cant open DB SSL Key file"
	ErrCantOpenSslRootCert string = "Cant open DB SSL Root Cert file"
	ErrNoDbConnection      string = "Cant connect to PostgreSQL Database"
)

type Config struct {
	Port          int           `json:"port"`
	HashCost      int           `json:"hashCost"`
	MigrationPath string        `json:"migrationPath"`
	Secret        string        `json:"secret"`
	Postgre       PostgreConfig `json:"postgre"`
}

type PostgreConfig struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	DbName      string `json:"dbName"`
	SslMode     string `json:"sslMode"`
	SslCert     string `json:"sslCert"`
	SslKey      string `json:"sslKey"`
	SslRootCert string `json:"sslRootCert"`
	Timeout     int    `json:"timeout"`
}

func (c Config) ValidateSslMode() error {
	validModes := []string{"disable", "require", "verify-ca", "verify-full"}
	check := false
	for _, validMode := range validModes {
		if c.Postgre.SslMode == validMode {
			check = true
		}
	}
	if !check {
		return errors.New("ErrInvalidSslMode")
	}
	return nil
}

func (c Config) ValidateDbTimeout() error {
	if c.Postgre.Timeout < 0 {
		return errors.New(ErrInvalidTimeout)
	}
	return nil
}

func (c Config) Validate() error {
	var errorList error
	if c.Postgre.SslMode != "disable" {
		_, err := os.OpenFile(c.Postgre.SslCert, os.O_RDONLY, 666)
		if err != nil {
			errorList = multierror.Append(errorList, errors.New(ErrCantOpenSslCert))
		}
		_, err = os.OpenFile(c.Postgre.SslKey, os.O_RDONLY, 666)
		if err != nil {
			errorList = multierror.Append(errorList, errors.New(ErrCantOpenSslKey))
		}

		_, err = os.OpenFile(c.Postgre.SslRootCert, os.O_RDONLY, 666)
		if err != nil {
			errorList = multierror.Append(errorList, errors.New(ErrCantOpenSslRootCert))
		}
	}

	connectionString := fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v sslmode=%v sslcert=%v sslkey=%v sslrootcert=%v connect_timeout=%v", c.Postgre.DbName, c.Postgre.Username, c.Postgre.Password, c.Postgre.Host, c.Postgre.Port, c.Postgre.SslMode, c.Postgre.SslCert, c.Postgre.SslKey, c.Postgre.SslRootCert, c.Postgre.Timeout)

	db, err := sql.Open("postgres", connectionString)
	defer db.Close()
	if err != nil {
		errorList = multierror.Append(errorList, errors.New(ErrNoDbConnection))
	}
	err = db.Ping()
	if err != nil {
		errorList = multierror.Append(errorList, errors.New(ErrNoDbConnection))
	}
	return errorList
}
