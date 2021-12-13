package main

import (
	"io/ioutil"
	"os"
	"reloaded"
	"testing"
)

func TestGoR(t *testing.T) {

	// Set variables for arguments

	file_name := "file_name"
	inputFile := "sample.txt"
	outputFile := "output.txt"

	args := []string{file_name, inputFile, outputFile}

	// Create a test table with input and output results

	tables := []struct {
		input  string
		output string
	}{
		{
			"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			"If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"},

		{
			"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			"I have to pack 5 outfits. Packed 26 just to be sure"},

		{
			"Don't be sad ,because sad backwards is das . And das not good",
			"Don't be sad, because sad backwards is das. And das not good",
		},
		{
			"harold wilson (cap, 2) : ' I’m a optimist ,but a optimist who carries a raincoat . '",
			"Harold Wilson: 'I’m an optimist, but an optimist who carries a raincoat.'",
		},
		{
			"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			"It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			"Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			"There is no greater agony than bearing a untold story inside you.",
			"There is no greater agony than bearing an untold story inside you.",
		},
		{
			"Punctuation tests are ... kinda boring ,don't you think !?",
			"Punctuation tests are... kinda boring, don't you think!?",
		},
	}

	for _, test_list := range tables {

		// Write the input strings to a file as a slice of bytes

		if err := ioutil.WriteFile(inputFile, []byte(test_list.input), os.ModePerm); err != nil {
			t.Errorf("Could not write to file")
		}

		// Run the function
		os.Args = args
		reloaded.GoReloaded()

		// Read the results of the output file and compare the input to the output
		if result, err := ioutil.ReadFile(outputFile); err != nil {
			t.Errorf("Could not read file")
		} else if string(result) != test_list.output {
			t.Errorf("The output of Go reloaded was incorrect, got: %s, want: %s.", string(result), test_list.output)
		}
	}
}
