package main

import (
	"simpledict/dict"

	"github.com/fatih/color"
)

func getAllDicts() *[]dict.Dict {
	return &[]dict.Dict{
		dict.NewCambridge(),
		dict.NewFreeDictionary(),
		dict.NewVocabulary(),
	}
}

func runDict(word *string, d *dict.Dict) (*dict.Result, error) {
	return (*d).GetDefinition(word)
}

func printResult(r *dict.Result, syn *bool) {
	{
		c := color.New(color.BgHiMagenta, color.Underline)
		c.Println("Definitions :")
		c = color.New(color.FgHiMagenta)
		for _, d := range r.Definition {
			if d == "" {
				continue
			}
			c.Printf("%s\n", d)
		}
	}

	if *syn {
		c := color.New(color.BgHiCyan, color.FgBlack, color.Underline)
		c.Println("Synonyms :")
		c = color.New(color.FgHiCyan)
		for _, s := range r.Synonyms {
			if s == "" {
				continue
			}
			c.Printf("%s\n", s)
		}
	}

}
