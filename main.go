package main

import (
	"flag"
	"fmt"
	"os"
	"simpledict/dict"
	"sync"

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

		// results := make(map[*string]*dict.Result) // store results used for final printing
		wg := new(sync.WaitGroup)
		for i := range *dicts {
			wg.Add(1)
			go func(d *dict.Dict) {
				defer wg.Done()
				r, err := runDict(&word, d)
				if err != nil {
					return
				}
				// results[(*d).GetName()] = r // final printing

				c := color.New(color.BgHiGreen, color.FgBlack, color.Underline)
				c.Printf("Source: %s\n", *(*d).GetName())
				printResult(r, syn)
				fmt.Println("")
			}(&(*dicts)[i])

		}
		wg.Wait()

		/* final printing
		for name, r := range results {
			c := color.New(color.BgHiGreen, color.FgBlack, color.Underline)
			c.Printf("Source: %s\n", *name)
			printResult(r, syn)
			fmt.Println("")
		}
		*/
	}
}
