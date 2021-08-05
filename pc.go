package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		// maybe print help here?
		os.Exit(1)
	}

	for _, arg := range args[1:] {
		fmt.Printf("%s ", escape(arg))
	}
	fmt.Println()

	cmd := exec.Command(
		args[1],
		args[2:]...,
	)
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

}

func escape(s string) string {
	if !strings.Contains(s, " ") {
		return s
	}

	if strings.Contains(s, "\"") {
		return fmt.Sprintf("'%s'", s)
	}

	return fmt.Sprintf("\"%s\"", s)
}
