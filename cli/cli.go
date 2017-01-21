package cli

import "github.com/urfave/cli"
import "github.com/NorbertKa/GoToken/config"

func NewGoToken() (*cli.App, *config.Config) {

	conf := config.Config{}

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "rHost",
			Value:       "localhost",
			EnvVar:      "GOTOKEN_RHOST",
			Destination: &conf.Redis.Host,
		},
		cli.IntFlag{
			Name:        "rPort",
			Value:       6379,
			EnvVar:      "GOTOKEN_RPORT",
			Destination: &conf.Redis.Port,
		},
		cli.StringFlag{
			Name:        "rPassword",
			Value:       "",
			EnvVar:      "GOTOKEN_RPASSWORD",
			Destination: &conf.Redis.Password,
		},
		cli.IntFlag{
			Name:        "rDb",
			Value:       0,
			EnvVar:      "GOTOKEN_RDB",
			Destination: &conf.Redis.DB,
		},
		cli.StringFlag{
			Name:        "pHost",
			Value:       "localhost",
			Usage:       "Host for PostgreSQL database",
			EnvVar:      "GOTOKEN_PHOST",
			Destination: &conf.Postgre.Host,
		},
		cli.IntFlag{
			Name:        "pPort",
			Value:       5432,
			Usage:       "Port for PostgreSQL database",
			EnvVar:      "GOTOKEN_PPORT",
			Destination: &conf.Postgre.Port,
		},
		cli.StringFlag{
			Name:        "pUsername",
			Value:       "",
			Usage:       "Username for PostgreSQL database",
			EnvVar:      "GOTOKEN_PUSERNAME",
			Destination: &conf.Postgre.Username,
		},
		cli.StringFlag{
			Name:        "pPassword",
			Value:       "",
			Usage:       "Password for PostgreSQL database",
			EnvVar:      "GOTOKEN_PPASSWORD",
			Destination: &conf.Postgre.Password,
		},
		cli.StringFlag{
			Name:        "pDbName",
			Value:       "",
			Usage:       "Name of the database to connect to",
			EnvVar:      "GOTOKEN_PDBNAME",
			Destination: &conf.Postgre.DbName,
		},
		cli.StringFlag{
			Name:        "pSslMode",
			Value:       "disable",
			Usage:       "SSL Mode for PostgreSQL database",
			EnvVar:      "GOTOKEN_PSSLMODE",
			Destination: &conf.Postgre.SslMode,
		},
		cli.StringFlag{
			Name:        "pSslCert",
			Value:       "",
			Usage:       "SSL Cert file location",
			EnvVar:      "GOTOKEN_PSSLCERT",
			Destination: &conf.Postgre.SslCert,
		},
		cli.StringFlag{
			Name:        "pSslKey",
			Value:       "",
			Usage:       "SSL Key file location",
			EnvVar:      "GOTOKEN_PSSLKEY",
			Destination: &conf.Postgre.SslKey,
		},
		cli.StringFlag{
			Name:        "pSslRootCert",
			Value:       "",
			Usage:       "SSL Root Cert file location",
			EnvVar:      "GOTOKEN_PSSLROOTCERT",
			Destination: &conf.Postgre.SslRootCert,
		},
		cli.IntFlag{
			Name:        "pTimeout",
			Value:       30,
			Usage:       "Maximum wait for connection, in seconds.",
			EnvVar:      "GOTOKEN_PTIMEOUT",
			Destination: &conf.Postgre.Timeout,
		},
		cli.IntFlag{
			Name:        "port",
			Value:       4812,
			Usage:       "Port",
			EnvVar:      "GOTOKEN_PORT",
			Destination: &conf.Port,
		},
		cli.IntFlag{
			Name:        "hashCost",
			Value:       16,
			Usage:       "Bcrypt hash cost",
			EnvVar:      "GOTOKEN_HASHCOST",
			Destination: &conf.HashCost,
		},
		cli.StringFlag{
			Name:        "migrationPath, mig",
			Value:       "./migrations",
			Usage:       "Migration path",
			EnvVar:      "GOTOKEN_MIGRATIONPATH",
			Destination: &conf.MigrationPath,
		},
		cli.StringFlag{
			Name:        "secret, s",
			Value:       "",
			Usage:       "Secret used to generate jwt tokens",
			EnvVar:      "GOTOKEN_SECRET",
			Destination: &conf.Secret,
		},
		cli.IntFlag{
			Name: "tokenDuration",
			Value: 15,
			Usage: "Regular token duration in minutes",
			EnvVar: "GOTOKEN_TOKENDURATION",
			Destination: &conf.TokenDuration,
		},
		cli.IntFlag{
			Name: "rTokenDuration",
			Value: 1440,
			Usage: "Refresh token duration in minutes",
			EnvVar: "GOTOKEN_RTOKENDURATION",
			Destination: &conf.RefreshTokenDuration,
		},
		cli.IntFlag{
			Name: "maxLogins",
			Value: 0,
			Usage: "Max logins from one user, 0 = Inf",
			EnvVar: "GOTOKEN_MAXLOGINS",
			Destination: &conf.MaxLogins,
		},
		cli.StringFlag{
			Name: "publicKey",
			Value: "",
			Usage: "Public Key",
			EnvVar: "GOTOKEN_PUBLICKEY",
			Destination: &conf.PublicKey,
		},
		cli.StringFlag{
			Name: "privateKey",
			Value: "",
			Usage: "Private Key",
			EnvVar: "GOTOKEN_PRIVATEKEY",
			Destination: &conf.PrivateKey,
		},
	}

	app.Name = "GoToken"
	app.Usage = "Auth service based on golang and JWT"

	return app, &conf
}
