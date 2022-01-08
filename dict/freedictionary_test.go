package dict

import "testing"

func TestFreeDictionaryImplementation(t *testing.T) {
	var _ Dict = (*FreeDictionary)(nil)
	var _ Dict = &FreeDictionary{}
	var _ Dict = new(FreeDictionary)
	t.Log("FreeDictionary successfully implements Dict interface")
}

func TestFreeDictionaryGetDefinition(t *testing.T) {
	word := "eating"
	d := new(FreeDictionary)
	results, err := d.GetDefinition(&word)
	if err != nil {
		t.Error(err)
	}

	if len(results.Definition) == 0 {
		t.Error("no definitions found")
	}
	if len(results.Synonyms) == 0 {
		t.Error("no synonyms found")
	}

	t.Log("FreeDictionary successfully returns definition")
	t.Logf("%#v", results)
}
