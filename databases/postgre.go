package database

import (
	"database/sql"
	"fmt"

	"github.com/NorbertKa/GoToken/config"
	_ "github.com/lib/pq"
)

type Postgre struct {
	*sql.DB
}

func NewPostgre(c *config.Config) (*Postgre, error) {
	connectionString := fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v sslmode=%v sslcert=%v sslkey=%v sslrootcert=%v connect_timeout=%v", c.Postgre.DbName, c.Postgre.Username, c.Postgre.Password, c.Postgre.Host, c.Postgre.Port, c.Postgre.SslMode, c.Postgre.SslCert, c.Postgre.SslKey, c.Postgre.SslRootCert, c.Postgre.Timeout)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	postgre := Postgre{db}
	return &postgre, nil
}
