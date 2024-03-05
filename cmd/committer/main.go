package main

import (
	"fmt"
	"os"

	"github.com/blmarquess/committer/internal/domain/entity"
	"github.com/blmarquess/committer/internal/git"
	"github.com/blmarquess/committer/internal/util"
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
	items := []util.Item{
		util.Item{}.New("✨ feat", "Add new feature"),
		util.Item{}.New("🐞 fix", "Fix a bug"),
		util.Item{}.New("🧪 test", "Add or update tests"),
		util.Item{}.New("♻️ refactor", "Code changes that neither fixes a bug nor adds a feature"),
		util.Item{}.New("🎨 style", "Code style changes (whitespace, formatting, etc.)"),
		util.Item{}.New("🚧 wip", "Work in progress"),
		util.Item{}.New("📚 docs", "Update documentation"),
		util.Item{}.New("📦 build", "Changes related to build process"),
		util.Item{}.New("♾️ ci", "Changes to CI configuration or scripts"),
		util.Item{}.New("⚡️ perf", "Performance improvements"),
		util.Item{}.New("↩ revert", "Reverts a previous commit"),
		util.Item{}.New("🔧 chore", "Changes to the build process, auxiliary tools, etc."),
		util.Item{}.New("🗑️ clean", "Removed unused code."),
	}

	var c entity.CommitEntity
	c.CommitType = util.GetListWithDesc(items)
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
		fmt.Println("Error on commit:", err)
		os.Exit(1)
	}
}
