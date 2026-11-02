// repl func... honestly not sure what the goal of it is yet

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func cleaninput(text string) []string {
	if text == "" {
		return []string{}
	}
	newSlice := []string{}
	newSlice = strings.Split(strings.ToLower(strings.TrimSpace(text)), " ")
	return newSlice
}

func startrepl(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	for {
		fmt.Fprint(writer, "Pokedex >")
		ok := scanner.Scan()
		if !ok {
			fmt.Fprintf(os.Stderr, "Writer closed")
			break
		}
		err := scanner.Err()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan Error: %v", err)
			break
		}
		input := scanner.Text()
		command := cleaninput(input)
		if len(command) == 0 {
			fmt.Fprintf(writer, "Please input a command\n")
			continue
		}
		if _, ok := CommandRegistery[command[0]]; !ok {
			fmt.Fprintf(writer, "please input a valid command\n")
			continue
		}
		CommandRegistery[command[0]].callback(writer)
	}
}
