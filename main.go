package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Args is too short!")
		os.Exit(0)
	}
	re := regexp.MustCompile(os.Args[1])
	var raw string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		raw += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
	output := re.ReplaceAllLiteralString(raw, os.Args[2])
	fmt.Println(output)
}
