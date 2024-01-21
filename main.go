package main

import (
	"commiter/modules"
	"fmt"
	"os"
)

func main() {
	stageEmpty, err := modules.IsStageEmpty()
	if err != nil {
		fmt.Println("Error on check stage:", err)
		os.Exit(1)
	}
	if stageEmpty {
		fmt.Println("Stage is empty, please add files.")
		os.Exit(1)
	}

	var commit modules.Commit
	commit.CommitType = modules.GetListWithDesc()
	commit.ShortDesc = modules.GetInput("Short Description: ", "")

	if modules.Condition("Add scope?", false) {
		commit.Scope = modules.GetInput("Scope of commit: ", "")
	}

	if modules.Condition("Add a long description?", false) {
		commit.LongDesc = modules.GetInput("Long description: ", "")
	}

	if modules.Condition("Add a id of issue? ", false) {
		commit.Issue = modules.GetInput("Issue number: ", "")
	}

	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J") // Clear screen
	if modules.GitCommit(modules.Formatter(commit)) != nil {
		fmt.Println("Erro ao commitar:", err)
		os.Exit(1)
	}
}
