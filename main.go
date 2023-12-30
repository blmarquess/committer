package main

import (
	"bufio"
	"commiter/modules"
	"fmt"
	"os"
	"strings"
)

func main() {
	options := modules.ListOptions{
		Items: []string{"feat", "fix", "docs", "style", "refactor", "test", "chore", "build", "ci", "perf", "revert"},
		Title: "Qual o tipo de commit?",
	}

	commitType, err := modules.CreateList(options)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	commit := getInput("Descrição curta: ", "")
	res := condition("Adicionar uma descrição longa no commit?", false)
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

	fmt.Printf("Qual o tipo de commit? %s\n", commitType)
	commitResult := "git commit -m \"" + "(" + commitType + "):" + commit + "\n\n" + longDescription + "\n\n " + issue + "\""
	fmt.Printf(commitResult)
}

func getInput(label string, def string) string {
	reader := bufio.NewReader(os.Stdin)
	var inputs string
	for {
		fmt.Fprintf(os.Stderr, "%s", label)
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
