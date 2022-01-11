package dict

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

type Vocabulary struct {
	*dict
}

func NewVocabulary() *Vocabulary {
	return &Vocabulary{
		dict: &dict{
			Name: "Vocabulary",
		},
	}
}

func (v *Vocabulary) GetDefinition(s *string) (*Result, error) {
	results := new(Result)
	// scrape using colly
	{
		c := colly.NewCollector()
		extensions.RandomUserAgent(c)
		extensions.Referer(c)
		c.IgnoreRobotsTxt = true

		c.OnHTML(".definitionsContainer .short", func(e *colly.HTMLElement) {
			// get definition
			results.Definition = append(results.Definition, e.Text)
		})

		c.OnHTML(".definitionsContainer .long", func(e *colly.HTMLElement) {
			// get definition
			results.Definition = append(results.Definition, e.Text)
		})

		c.OnHTML(".definitionsContainer ol > li > .definition", func(e *colly.HTMLElement) {
			el := e.DOM.Clone()
			el.Find(".pos-icon").Remove()
			// get definition
			txt := strings.TrimSpace(el.Text())
			results.Definition = append(results.Definition, txt)
		})

		c.OnHTML(".word-definitions .defContent .instances span > a.word", func(h *colly.HTMLElement) {
			// get synonyms
			results.Synonyms = append(results.Synonyms, h.Text)
		})
		if err := c.Visit(fmt.Sprintf("https://www.vocabulary.com/dictionary/%s", *s)); err != nil {
			return nil, err
		}
		c.Wait()
	}

	return results, nil
}
