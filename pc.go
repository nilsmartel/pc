package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 || args[1] == "--help" || args[1] == "-h" {
		printHelp()
		os.Exit(0)
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
		os.Exit(1)
	}

}

func printHelp() {
	fmt.Println("pc\n" +
		"	usage: 		pc <command> <...arguments>\n" +
		"	example:	pc mv lecture_01.pdf '01 Lecture.pdf'\n" +
		"\n" +
		"	Using pc you can print commands, before executing them.\n" +
		"	The arguments to pc are simply shell commands, along with their arguments.\n" +
		"\n" +
		"	This way you can easier document, what is happening behind the scenes.")
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
