package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Condition(label string, def bool) bool {
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
