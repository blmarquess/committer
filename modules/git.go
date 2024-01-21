package modules

import (
	"os/exec"
	"strings"
)

func IsStageEmpty() (bool, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}
	outputString := string(output)
	return strings.TrimSpace(outputString) == "", nil
}

func GetStagedFiles() ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", "--cached")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	outputString := string(output)
	lines := strings.Split(strings.TrimSpace(outputString), "\n")
	return lines, nil
}

func GetUntrackedFiles() ([]string, error) {
	cmd := exec.Command("git", "ls-files", "--others", "--exclude-standard")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	outputString := string(output)
	lines := strings.Split(strings.TrimSpace(outputString), "\n")
	return lines, nil
}

func GitCommit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func test() {
	// if len(untrackedFiles) > 0 {
	// 	fmt.Println("Arquivos não versionados:")
	// 	for _, file := range untrackedFiles {
	// 		fmt.Println(file)
	// 	}
	// }

	// stagedFiles, err := modules.GetStagedFiles()
	// if err != nil {
	// 	fmt.Println("Erro ao obter arquivos no stage:", err)
	// 	os.Exit(1)
	// }

	// if len(stagedFiles) > 0 {
	// 	fmt.Println("Arquivos no stage:")
	// 	for _, file := range stagedFiles {
	// 		fmt.Println(file)
	// 	}
	// } else {
	// 	fmt.Println("Nenhum arquivo no stage encontrado.")
	// }
}
