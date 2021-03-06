package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	builtins = []string{"cd", "help", "exit"}
	flist    = []func([]string) int{eshell_cd, eshell_help, eshell_exit}
)

func main() {
	eshell_loop()
}

func eshell_readline() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func eshell_split_line(line string) []string {
	tokens := strings.Split(line, " ")

	return tokens
}

func eshell_execute(args []string) int {
	if args[0] == "" {
		return 1
	}

	// check for builtins
	for i, builtCmd := range builtins {
		if strings.Compare(args[0], builtCmd) == 0 {
			funcStat := flist[i]
			return funcStat(args)
		}
	}

	return eshell_launch(args)
}

func eshell_launch(args []string) int {
	tempArgs := args[1:]
	cmd := exec.Command(args[0], tempArgs...)
	cmd.Env = os.Environ()
	var out bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stdErr
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(out.String())
	fmt.Print(stdErr.String())
	return 1
}

func eshell_loop() {
	status := 1

	for status != 0 {
		fmt.Print("> ")
		line := eshell_readline()
		tokens := eshell_split_line(line)
		status = eshell_execute(tokens)
	}
}
