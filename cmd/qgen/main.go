package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mizumoto-cn/leetcodego/cmd/qgen/pkg"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "qgen",
		Usage: "generate leetcode solvation from template",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "target",
				Aliases: []string{"t"},
				Value:   "./",
				Usage:   "target folder path",
			},
			&cli.StringFlag{
				Name:    "template",
				Aliases: []string{"f"},
				Value:   "",
				Usage:   "template path",
			},
			&cli.StringFlag{
				Name:    "name_or_umber",
				Aliases: []string{"n"},
				Value:   "NewProblem",
				Usage:   "Problem name or number, like leetcode problem numbers",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Name:", c.String("name_or_umber"))
			err := pkg.Generate(c.String("target"), c.String("template"), c.String("name_or_umber"))
			return err
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
