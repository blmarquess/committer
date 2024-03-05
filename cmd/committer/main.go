package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Erro ao obter o diretório home:", err)
		return
	}

	jsonFilePath := filepath.Join(homeDir, ".commiter", "config.json")
	file, err := os.Open(jsonFilePath)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	var config entity.Configuration
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}

	var selectSteps []entity.Step
	for _, step := range config.Steps {
		if step.Type == "select" {
			selectSteps = append(selectSteps, step)
		}
	}
	var items []util.Item
	for _, step := range selectSteps {
		for _, option := range step.Options {
			items = append(items, util.Item{}.New(option.Value, option.Message))
		}
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
