package main

import (
	"fmt"
	"log"

	"github.com/esoytekin/utility"
	"github.com/urfave/cli"
)

func testCreateApp() {
	auth := cli.Author{
		Name:  "Emrah Soytekin",
		Email: "emrahsoytekin@gmail.com",
	}

	app := utility.CreateNewApp("Dovetail", "automated tasks", auth)

	c := app.AddCommand("test", "test command", []string{"t"})

	c.Action = func(c *cli.Context) error {
		fmt.Println("working")
		return nil
	}

	app.AddCommand("test2", "test2 usage", []string{"t2"})

	action := func(c *cli.Context) error {
		return cli.ShowAppHelp(c)
	}
	app.SetAction(action)

	err := app.Start()

	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	testCreateApp()
}
