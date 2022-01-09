package dict

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type FreeDictionary struct {
	*dict
}

func NewFreeDictionary() *FreeDictionary {
	return &FreeDictionary{
		dict: &dict{
			Name: "FreeDictionary",
		},
	}
}

func (d *FreeDictionary) GetDefinition(s *string) (*Result, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", *s), nil)
	if err != nil {
		return nil, err
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	type result struct {
		Word      string `json:"word"`
		Phonetics []struct {
			Text  string `json:"text"`
			Audio string `json:"audio"`
		} `json:"phonetics"`
		Meaning []struct {
			PartOfSpeech string `json:"partOfSpeech"`
			Definition   []struct {
				Definition string   `json:"definition"`
				Example    string   `json:"example"`
				Synonyms   []string `json:"synonyms"`
			} `json:"definitions"`
		} `json:"meanings"`
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("failed to contact freeDictionary API")
	}

	results := make([]result, 0)
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, errors.New("no definition results found")
	}

	finalresult := new(Result)

	for _, r := range results {
		for _, m := range r.Meaning {
			for _, d := range m.Definition {
				finalresult.Definition = append(finalresult.Definition, d.Definition)
				finalresult.Synonyms = append(finalresult.Synonyms, d.Synonyms...)
			}
		}
	}

	return finalresult, nil
}
