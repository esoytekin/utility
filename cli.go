package utility

import (
	"os"
	"time"

	"github.com/urfave/cli"
)

type AppTemplate struct {
	app *cli.App
}

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

func (self *AppTemplate) AddCommand(name, usage string, aliases []string) *cli.Command {

	var c = cli.Command{Name: name, Usage: usage, Aliases: aliases}

	if self.app.Commands == nil {
		self.app.Commands = []cli.Command{}
	}

	self.app.Commands = append(self.app.Commands, c)

	l := len(self.app.Commands)

	return &self.app.Commands[l-1]

}

func (self AppTemplate) SetAction(action func(c *cli.Context) error) {

	self.app.Action = action

}

func (self AppTemplate) Start() error {

	err := self.app.Run(os.Args)

	if err != nil {
		return err
	}

	return nil
}
