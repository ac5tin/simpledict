package main

import (
	"fmt"
	"simpledict/dict"
)

func getAllDicts() *[]dict.Dict {
	return &[]dict.Dict{
		new(dict.Cambridge),
		new(dict.FreeDictionary),
	}
}

func runDict(word *string, d *dict.Dict) (*dict.Result, error) {
	return (*d).GetDefinition(word)
}

func printResult(r *dict.Result) {
	fmt.Println("Definitions :")
	for _, d := range r.Definition {
		fmt.Printf("%s\n", d)
	}
	fmt.Println("Synonyms :")
	for _, s := range r.Synonyms {
		fmt.Printf("%s\n", s)
	}
}