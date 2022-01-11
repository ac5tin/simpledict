package dict

import "testing"

func TestVocabularyImplementation(t *testing.T) {
	var _ Dict = (*Vocabulary)(nil)
	var _ Dict = &Vocabulary{}
	var _ Dict = new(Vocabulary)

	v := NewVocabulary()
	if v.GetName() == nil {
		t.Error("GetName() should not return nil")
	}

	if *v.GetName() != "Vocabulary" {
		t.Errorf("expected Vocabulary.GetName() to return 'Vocabulary', got %s", *v.GetName())
	}

	t.Log("Vocabulary successfully implements Dict interface")
}

func TestVocabularyGetDefinition(t *testing.T) {
	word := "imbue"
	d := new(Vocabulary)
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

	t.Log("Vocabulary successfully returns definition")
	t.Logf("%#v", results)
}
