package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Shell struct {
	builtinCmds map[CmdIdent]CmdHandler
}

func NewShell() *Shell {
	shell := &Shell{
		builtinCmds: make(map[CmdIdent]CmdHandler),
	}
	shell.registerBuiltinCmd(CmdExit, shell.exitCmd)
	shell.registerBuiltinCmd(CmdEcho, shell.echoCmd)
	shell.registerBuiltinCmd(CmdType, shell.typeCmd)
	shell.registerBuiltinCmd(CmdPwd, shell.pwdCmd)
	shell.registerBuiltinCmd(CmdCd, shell.cdCmd)
	return shell
}

func (s *Shell) registerBuiltinCmd(cmd CmdIdent, handler CmdHandler) {
	s.builtinCmds[cmd] = handler
}

type CmdIdent string

type CmdHandler func(args ...string)

const (
	CmdExit CmdIdent = "exit"
	CmdEcho CmdIdent = "echo"
	CmdType CmdIdent = "type"
	CmdPwd  CmdIdent = "pwd"
	CmdCd   CmdIdent = "cd"
)

func (s *Shell) exitCmd(args ...string) {
	_ = args
	os.Exit(0)
}

func (s *Shell) echoCmd(args ...string) {
	fmt.Println(strings.Join(args, " "))
}

func (s *Shell) typeCmd(args ...string) {
	command := args[0]
	if _, ok := s.builtinCmds[CmdIdent(command)]; ok {
		fmt.Printf("%s is a shell builtin\n", command)
		return
	}
	if path, err := exec.LookPath(command); err == nil {
		fmt.Printf("%s is %s\n", command, path)
		return
	}
	fmt.Printf("%s: not found\n", command)
}

func (s *Shell) pwdCmd(args ...string) {
	_ = args
	wd, _ := os.Getwd()
	fmt.Println(wd)
}

func (s *Shell) cdCmd(args ...string) {
	dir := args[0]
	if dir == "~" {
		dir = os.Getenv("HOME")
	}
	if err := os.Chdir(dir); err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", dir)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	shell := NewShell()

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

		if handler, ok := shell.builtinCmds[CmdIdent(command)]; ok {
			handler(args...)
			continue
		}

		if _, err := exec.LookPath(command); err == nil {
			exCmd := exec.Command(command, args...)
			exCmd.Stdout = os.Stdout
			exCmd.Stderr = os.Stderr
			_ = exCmd.Run()
			continue
		}

		io.WriteString(
			os.Stdout,
			fmt.Sprintf("%s: command not found\n", command))

	}
}
