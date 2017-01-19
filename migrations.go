package main

import (
	"fmt"

	"github.com/NorbertKa/GoToken/config"
	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/mattes/migrate/migrate"
)

func MigrateUp(conf *config.Config) (bool, []error) {
	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v?sslmode=%v&connect_timeout=%v&sslcert=%v&sslkey=%v&sslrootcert=%v",
		conf.Postgre.Username,
		conf.Postgre.Password,
		conf.Postgre.Host,
		conf.Postgre.Port,
		conf.Postgre.SslMode,
		conf.Postgre.Timeout,
		conf.Postgre.SslCert,
		conf.Postgre.SslKey,
		conf.Postgre.SslRootCert)
	allErrors, ok := migrate.UpSync(connectionString, conf.MigrationPath)
	return ok, allErrors
}

func MigrateDown(conf *config.Config) (bool, []error) {
	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v?sslmode=%v&connect_timeout=%v&sslcert=%v&sslkey=%v&sslrootcert=%v",
		conf.Postgre.Username,
		conf.Postgre.Password,
		conf.Postgre.Host,
		conf.Postgre.Port,
		conf.Postgre.SslMode,
		conf.Postgre.Timeout,
		conf.Postgre.SslCert,
		conf.Postgre.SslKey,
		conf.Postgre.SslRootCert)
	allErrors, ok := migrate.DownSync(connectionString, conf.MigrationPath)
	return ok, allErrors
}
