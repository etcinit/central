package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/etcinit/central/app"
	"github.com/etcinit/central/database"
	"github.com/facebookgo/inject"
	"github.com/jacobstr/confer"
	"github.com/kr/pretty"
)

func main() {
	// Create the configuration
	// In this case, we will be using the environment and some safe defaults
	config := confer.NewConfig()
	config.SetDefault("database.driver", "sqlite")
	config.SetDefault("database.file", ":memory:")
	config.ReadPaths("config/main.yml", "config/main.production.yml")
	config.AutomaticEnv()

	pretty.Println(config.AllSettings())

	// Next, we setup the dependency graph
	// In this example, the graph won't have many nodes, but on more complex
	// applications it becomes more useful.
	var g inject.Graph
	var central app.Central
	g.Provide(
		&inject.Object{Value: config},
		&inject.Object{Value: &central},
		&inject.Object{Value: &database.Directory{}},
	)
	g.Populate()

	// Setup the command line application
	app := cli.NewApp()
	app.Name = "central"
	app.Usage = "Central Configuration Service"

	// Set version and authorship info
	app.Version = "0.0.3"
	app.Author = "Eduardo Trujillo <ed@chromabits.com>"

	// Setup the default action. This action will be triggered when no
	// subcommand is provided as an argument
	app.Action = func(c *cli.Context) {
		fmt.Println("Usage: central [global options] command [command options] [arguments...]")
	}

	app.Commands = []cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"s", "server", "listen"},
			Usage:   "starts the API server",
			Action:  central.Serve.Run,
		},
	}

	// Begin
	app.Run(os.Args)
}
