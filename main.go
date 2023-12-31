package main

import (
	"commiter/modules"
	"fmt"
	"os"
)

func main() {
	// options := modules.ListOptions{
	// 	Items: []modules.Item{
	// 		modules.Item{Title: "feat", Desc: "nova funcionalidade"},
	// 		modules.Item{Title: "fix", Desc: "descrição"},
	// 		modules.Item{Title: "docs", Desc: "descrição"},
	// 		modules.Item{Title: "style", Desc: "descrição"},
	// 		modules.Item{Title: "refactor", Desc: "descrição"},
	// 		modules.Item{Title: "test", Desc: "descrição"},
	// 		modules.Item{Title: "chore", Desc: "descrição"},
	// 		modules.Item{Title: "build", Desc: "descrição"},
	// 		modules.Item{Title: "ci", Desc: "descrição"},
	// 		modules.Item{Title: "perf", Desc: "descrição"},
	// 		modules.Item{Title: "revert", Desc: "descrição"},
	// 	},
	// 	Title: "Qual o tipo de commit?",
	// }
	options := modules.ListOptions{
		Items: []string{"feat", "fix", "docs", "style", "refactor", "test", "chore", "build", "ci", "perf", "revert"},
		Title: "Qual o tipo de commit?",
	}
	var commit string

	commitType, err := modules.CreateList(options)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	commit = commit + commitType
	commit = commit + modules.GetInput("Descrição curta: ", "")

	if modules.Condition("Adicionar uma descrição longa no commit?", false) {
		commit = commit + "\n\n" + modules.GetInput("Descrição longa", "")
	}

	if modules.Condition("Adicionar uma id de alguma issue ?", false) {
		commit = commit + "\n\n" + "#" + modules.GetInput("Issue: ", "")
	}

	fmt.Println(commit)

	// res := modules.GetListWithDesc()
	// fmt.Println(res)
}
