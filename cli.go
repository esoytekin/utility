package utility

import (
	"os"
	"time"

	"github.com/urfave/cli"
)

// AppTemplate is a wrapper for cli.App
// it adds some functionality like AddCommand, AddBoolFlag etc.
type AppTemplate struct {
	app *cli.App
}

// CreateNewApp creates AppTemplate with appname, usage and author
func CreateNewApp(appname, usage string, author cli.Author) *AppTemplate {
	app := cli.NewApp()
	app.Name = appname
	app.Usage = usage
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		author,
	}
	app.Copyright = "(c) 2019 Serious Enterprise"
	app.Version = "1.0.0"
	app.UseShortOptionHandling = true
	app.EnableBashCompletion = true

	return &AppTemplate{app}
}

// AddCommand adds new command with name, usage and aliases
func (templ *AppTemplate) AddCommand(name, usage string, aliases []string) *cli.Command {

	var c = cli.Command{Name: name, Usage: usage, Aliases: aliases}

	if templ.app.Commands == nil {
		templ.app.Commands = []cli.Command{}
	}

	templ.app.Commands = append(templ.app.Commands, c)

	l := len(templ.app.Commands)

	return &templ.app.Commands[l-1]

}

// AddStringFlag add string flag with name and usage
func (templ *AppTemplate) AddStringFlag(name, usage string) *cli.Flag {

	return addFlag(templ, true, name, usage)

}

// AddBoolFlag add bool flag with name and usage
func (templ *AppTemplate) AddBoolFlag(name, usage string) *cli.Flag {
	return addFlag(templ, false, name, usage)
}

func addFlag(temp *AppTemplate, isString bool, name, usage string) *cli.Flag {

	var f cli.Flag

	if isString {
		f = cli.StringFlag{Name: name, Usage: usage}
	} else {
		f = cli.BoolFlag{Name: name, Usage: usage}
	}

	temp.app.Flags = append(temp.app.Flags, f)

	l := len(temp.app.Flags)

	return &temp.app.Flags[l-1]

}

// GetCommands returns command list
func (templ AppTemplate) GetCommands() []cli.Command {

	return templ.app.Commands
}

// SetAction sets action for app
func (templ AppTemplate) SetAction(action func(c *cli.Context) error) {

	templ.app.Action = action
}

// Start starts cli app
func (templ AppTemplate) Start() error {

	err := templ.app.Run(os.Args)

	if err != nil {
		return err
	}

	return nil
}
