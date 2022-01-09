package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
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
			fmt.Println("==========================================================")
			fmt.Printf("Source: %s\n", *d.GetName())
			printResult(r, syn)
			fmt.Println("==========================================================")
		}
	}
}
