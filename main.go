package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func printUsage() {
	fmt.Println("Usage: simpledict [flags] <word>")
}

func main() {
	args := os.Args
	if len(args) < 2 {
		printUsage()
		os.Exit(1)
	}
	word := args[len(args)-1]

	// flags
	syn := flag.Bool("syn", false, "show synonyms")
	flag.Parse()

	// run all dicts
	{
		dicts := getAllDicts()
		for _, d := range *dicts {
			r, err := runDict(&word, &d)
			if err != nil {
				continue
			}
			c := color.New(color.BgHiGreen, color.FgBlack, color.Underline)
			c.Printf("Source: %s\n", *d.GetName())
			printResult(r, syn)
			fmt.Println("")
		}
	}
}
