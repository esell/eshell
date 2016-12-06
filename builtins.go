package main

import (
	"errors"
	"fmt"
	"os"
)

func eshell_cd(args []string) int {
	if args[1] == "" {
		fmt.Println(errors.New("lsh: expected argument to \"cd\""))
	} else {
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Println(err)
		}
	}
	return 1
}

func eshell_help(args []string) int {
	fmt.Println("eshell, a shitty shell")
	return 1
}

func eshell_exit(args []string) int {
	return 0
}
