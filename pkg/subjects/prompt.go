package subjects

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, reader *bufio.Reader) (string) {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func PromptOptions() {
	reader := bufio.NewReader(os.Stdin)
	opt := getInput("Choose option (a - add item, s - save album)", reader)

	fmt.Println(opt)
}