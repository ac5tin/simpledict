package dict

import "testing"

func TestCambridgeImplementation(t *testing.T) {
	var _ Dict = (*Cambridge)(nil)
	var _ Dict = &Cambridge{}
	var _ Dict = new(Cambridge)

	cb := NewCambridge()
	if cb.GetName() == nil {
		t.Error("GetName() should not return nil")
	}

	if *cb.GetName() != "Cambridge" {
		t.Errorf("expected Cambridge.GetName() to return 'Cambridge', got %s", *cb.GetName())
	}

	t.Log("Cambridge successfully implements Dict interface")
}

func TestCambridgeGetDefinition(t *testing.T) {
	word := "eating"
	d := new(Cambridge)
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

	t.Log("Cambridge successfully returns definition")
	t.Logf("%#v", results)
}
