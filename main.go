package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/asvvvad/clap/clap"
	"github.com/asvvvad/clap/clap/b"
	"github.com/atotto/clipboard"
)

func main() {
	var emoji string
	var file string
	var stdin bool
	var bf bool
	var print bool
	var noOutput bool

	flag.StringVar(&emoji, "emoji", "👏️", "Set emoji to use.")
	flag.StringVar(&file, "file", "", "Use text from file")
	flag.BoolVar(&stdin, "stdin", false, "Use text from stdin")
	flag.BoolVar(&bf, "b", false, "Replace the letter B with 🅱️ emoji")
	flag.BoolVar(&print, "print", false, "Print result.")
	flag.BoolVar(&noOutput, "no-out", false, "No output on success; use when starting it with a system shortcut")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println("C👏️L👏️A👏️P by asvvvad (c)lap 2020 - https://github.com/asvvvad/clap")
	}

	flag.Parse()

	// The string that will contain the input text
	var text string
	var err error
	var bytes []byte

	if stdin {
		bytes, err = ioutil.ReadAll(os.Stdin)
		text = string(bytes)
	} else if file != "" {
		bytes, err = ioutil.ReadFile(file)
		text = string(bytes)
	} else {
		text, err = clipboard.ReadAll()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// The string that will contain the final result
	result := ""

	if emoji != "" {
		result = clap.Emoji(text, emoji)
	} else {
		result = clap.Clap(text)
	}

	if bf {
		result = b.B(result)
	}

	if print {
		fmt.Println(result)
	} else {
		err := clipboard.WriteAll(result)
		if err != nil {
			fmt.Println(err)
			// Print out in case of unsupported system
			fmt.Println(result)
			os.Exit(1)
		}
		if !noOutput {
			fmt.Println("Clapped👏️and👏️copied!")
		}
	}
}
