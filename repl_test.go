// repl testing file

package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCleanInput(t *testing.T) {
	type testCase struct {
		input    string
		expected []string
	}

	cases := []testCase{
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "PiKaChU",
			expected: []string{"pikachu"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "Bulbasaur Blastois CHARIZARD",
			expected: []string{"bulbasaur", "blastois", "charizard"},
		},
	}
	for i, test := range cases {
		funcResult := cleaninput(test.input)
		if len(funcResult) != len(test.expected) {
			t.Fatalf("test: %d, Expected len: %d, Got: %d", i, len(test.expected), len(funcResult))
		}
		for j := 0; j < len(funcResult); j++ {
			if funcResult[j] != test.expected[j] {
				t.Fatalf("test: %d: Expected: %s, Got: %s",
					i, test.expected[j], funcResult[j])
			}
		}
	}
}

func TestStartRepl(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
	}
	cases := []testCase{
		{
			name:  "help",
			input: "help",
			expected: "Pokedex >Welcome to the Pokedex!\n" +
				"Usage:\n" +
				"help: Displays a help message\n" +
				"exit: Exits the Pokedex\n" +
				"Pokedex >",
		},
		{
			name:     "empty",
			input:    "\n",
			expected: "Pokedex >Please input a command\nPokedex >",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			reader := strings.NewReader(c.input)
			var writer bytes.Buffer
			startrepl(reader, &writer)
			gotWriter := writer.String()
			if gotWriter != c.expected {
				t.Errorf("\nExpected: %q, Got: %q", c.expected, gotWriter)
			}
		})
	}
}
