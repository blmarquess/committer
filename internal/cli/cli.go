package cli

import (
    "fmt"
    "log"
    "os"

    "github.com/urfave/cli/v2"
)

func CLI() {
	app := &cli.App{
		Name:  "Git Automation",
		Usage: "Automatiza comandos Git",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "pr",
				Usage: "Executa GitPR",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("pr") {
				success, err := GitPR()
				if err != nil {
					return cli.Exit(fmt.Sprintf("Erro: %s", err), 1)
				}
				if success {
					fmt.Println("Push realizado com sucesso!")
				}
			} else {
				success, err := GitCommit()
				if err != nil {
					return cli.Exit(fmt.Sprintf("Erro: %s", err), 1)
				}
				if success {
					fmt.Println("Commit realizado com sucesso!")
				}
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}