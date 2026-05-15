package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Print("$ ")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		input := scanner.Text()

		command := strings.Split(input, " ")[0]

		io.WriteString(
			os.Stdout,
			fmt.Sprintf("%s: command not found\n", command))
	}
}
