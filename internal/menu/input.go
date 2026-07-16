package menu

import (
	"bufio"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func Input() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}