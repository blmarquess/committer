package main

import (
	"fmt"
	"os"

	"github.com/blmarquess/commiter/internal/domain/entity"
	"github.com/blmarquess/commiter/internal/git"
	"github.com/blmarquess/commiter/internal/util"
)

func main() {
	stageEmpty, err := git.IsStageEmpty()
	if err != nil {
		fmt.Println("Error on check stage:", err)
		os.Exit(1)
	}
	if stageEmpty {
		fmt.Println("Stage is empty, please add files.")
		os.Exit(1)
	}

	var c entity.CommitEntity
	c.CommitType = util.GetListWithDesc()
	c.ShortDesc = util.GetInput("Short Description: ", "")

	if util.Condition("Add scope?", false) {
		c.Scope = util.GetInput("Scope of commit: ", "")
	}

	if util.Condition("Add a long description?", false) {
		c.LongDesc = util.GetInput("Long description: ", "")
	}

	if util.Condition("Add a id of issue? ", false) {
		c.Issue = util.GetInput("Issue number: ", "")
	}

	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J") // Clear screen
	if git.Commit(c.Formatter()) != nil {
		fmt.Println("Erro ao commitar:", err)
		os.Exit(1)
	}
}
