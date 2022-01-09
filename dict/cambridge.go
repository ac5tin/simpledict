package dict

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

type Cambridge struct {
	*dict
}

func (d *Cambridge) GetDefinition(s *string) (*Result, error) {
	results := new(Result)
	// scrape using colly
	{
		c := colly.NewCollector()
		extensions.RandomUserAgent(c)
		extensions.Referer(c)
		c.IgnoreRobotsTxt = true

		c.OnHTML("div.def.ddef_d", func(e *colly.HTMLElement) {
			// get definition
			results.Definition = append(results.Definition, e.Text)
		})
		c.OnHTML(".hul-u > li.had.t-i > a", func(h *colly.HTMLElement) {
			// get synonyms
			results.Synonyms = append(results.Synonyms, h.Text)
		})
		if err := c.Visit(fmt.Sprintf("https://dictionary.cambridge.org/dictionary/english/%s", *s)); err != nil {
			return nil, err
		}
		c.Wait()
	}
	return results, nil
}
