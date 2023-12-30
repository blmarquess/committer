package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/blmarques/committer/modules"
)

func main() {
	commitType := modules.CreateList("Qual o tipo de commit?", []string{"feat", "fix", "docs", "style", "refactor", "test", "chore", "build", "ci", "perf", "revert"})
	commit := getInput("Descrição curta", "")
	res := condition("Adicionar uma descrição longa no commit?", true)
	var longDescription string
	if res {
		longDescription = getInput("Descrição longa", "")
	}
	res = condition("Adicionar uma issue no commit?", false)
	var issue string
	if res {
		issue = getInput("Issue", "")
	}
	// return modules.Commit(commitType, commit, longDescription, issue)
	commitResult := "git commit -m \"" + "(" + commitType + "):" + commit + "\n\n" + longDescription + "\n\n " + issue + "\""
	fmt.Printf(commitResult)
}

func getInput(label string, def string) string {
	reader := bufio.NewReader(os.Stdin)
	var inputs string
	for {
		fmt.Fprintf(os.Stderr, "%s (%s) ", label, def)
		inputs, _ = reader.ReadString('\n')
		inputs = strings.TrimSpace(inputs)
		if inputs == "" {
			return def
		}
		return inputs
	}
}

func condition(label string, def bool) bool {
	choices := "Y/n"
	if !def {
		choices = "y/N"
	}

	reader := bufio.NewReader(os.Stdin)
	var inputs string
	for {
		fmt.Fprintf(os.Stderr, "%s (%s) ", label, choices)
		inputs, _ = reader.ReadString('\n')
		inputs = strings.TrimSpace(inputs)
		if inputs == "" {
			return def
		}
		inputs = strings.ToLower(inputs)
		if inputs == "y" || inputs == "yes" || inputs == "s" || inputs == "sim" {
			return true
		}
		if inputs == "n" || inputs == "no" || inputs == "não" || inputs == "nao" {
			return false
		}
	}
}
