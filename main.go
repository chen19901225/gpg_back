package main



import (
	"gpg_back/pkg/runner"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// var nameRaw, excludeName string
	var database, bin, dir string
	var backup = 0
	// var showDetail = 0
	var verbose = 1
	app := &cli.App{
		Name:    "gpg_back",
		Usage:   "gpg_back",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "database",
				Usage:       "backup database name",
				Required:    true,
				Destination: &database,
			},
			&cli.StringFlag{
				Name:        "bin",
				Usage:       "bin path",
				Required:    false,
				Value: "/www/server/pgsql/bin/pg_dump",
				DefaultText: "/www/server/pgsql/bin/pg_dump",
				Destination: &bin,
			},
			&cli.StringFlag{
				Name:        "dir",
				Usage:       "backup dir",
				Required:    false,
				Value: "/www/backup/pgsql_bak",
				DefaultText: "/www/backup/pgsql_bak",
				Destination: &dir,
			},
			&cli.IntFlag{
				Name:        "backup",
				Usage:       "backup count",
				Required:    false,
				Value: 7,
				DefaultText: "7",
				Destination: &backup,
			},
			&cli.IntFlag{
				Name:        "verbose",
				Usage:       "show log detail",
				Required:    false,
				Value:       1,
				DefaultText: "1",
				Destination: &verbose,
			},
		},
		Action: func(c *cli.Context) error {
			// log.Printf("开始")
			err := runner.Run(database, bin, dir, backup, verbose)
			if err != nil {
				log.Fatal(err)
			}
			return err
			// return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}

}
