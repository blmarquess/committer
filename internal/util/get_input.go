package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput(label string, def string) string {
	reader := bufio.NewReader(os.Stdin)
	var input string
	for {
		fmt.Fprintf(os.Stderr, "%s", label)
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			return def
		}
		return input
	}
}
