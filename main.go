package main

import "fmt"
import "github.com/urfave/cli"
import "os"

func setupCLIApp(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Initializes dot syncer",
			Action: func(c *cli.Context) error {
				if args := c.Args(); len(args) > 0 {
					fmt.Println("added task: ", args.First())
					return nil
				}
				//TODO return error and write help
			},
		},
		{
			Name:    "sync",
			Aliases: []string{"c"},
			Usage:   "Sync dot fileseee",
			Action: func(c *cli.Context) error {

				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
	}
	app.Name = "boom"
	app.Usage = "make an explosive entrance"

}
func main() {
	app := cli.NewApp()
	setupCLIApp(app)
	app.Run(os.Args)
}
