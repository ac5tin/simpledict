package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		os.Exit(1)
	}
	word := args[1]

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
			printResult(r)
			fmt.Println("==========================================================")
		}
	}
}
