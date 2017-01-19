package main

import (
	"fmt"
	"os"

	gcli "github.com/NorbertKa/GoToken/cli"
	"github.com/urfave/cli"
)

const version string = "0.0.1"

const (
	ErrUndefinedError string = "UNDEFINED ERROR"
)

func main() {
	app, conf := gcli.NewGoToken()
	if app == nil || conf == nil {
		panic(ErrUndefinedError)
	}
	app.Version = version

	app.Action = func(c *cli.Context) error {
		fmt.Println("#==> Validating Database")
		err := conf.ValidateDbTimeout()
		if err != nil {
			panic(err)
		}
		err = conf.ValidateSslMode()
		if err != nil {
			panic(err)
		}
		err = conf.ValidatePostgreConnection()
		if err != nil {
			panic(err)
		}
		fmt.Println("#==> DONE")
		fmt.Println("#==> Running DB Migrations")
		check, errors := MigrateUp(conf)
		if !check {
			for _, error := range errors {
				fmt.Println(error)
			}
			panic("DB Migration Error")
		}
		fmt.Println("#==> DONE")
		return nil
	}

	app.Run(os.Args)
}
