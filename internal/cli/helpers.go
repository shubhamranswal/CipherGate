package cli

import (
	"bufio"
	"os"
	"strings"
)

func readInput(
	prompt string,
) (string, error) {

	print(prompt)

	reader := bufio.NewReader(
		os.Stdin,
	)

	input, err := reader.ReadString(
		'\n',
	)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(
		input,
	), nil
}
