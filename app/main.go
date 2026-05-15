package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		input := scanner.Text()

		parts := strings.Split(input, " ")
		command := parts[0]
		args := parts[1:]

		if command == "exit" {
			break
		} else if command == "echo" {
			fmt.Println(strings.Join(args, " "))
			continue
		}

		io.WriteString(
			os.Stdout,
			fmt.Sprintf("%s: command not found\n", command))
	}
}
