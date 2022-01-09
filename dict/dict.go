package dict

import "errors"

type Result struct {
	Definition []string
	Synonyms   []string
}

type Dict interface {
	GetDefinition(s *string) (*Result, error)
	GetName() *string
}

type dict struct {
	Dict
	Name string
}

func (d *dict) GetDefinition(s *string) (*Result, error) {
	return nil, errors.New("not implemented")
}

func (d *dict) GetName() *string {
	return &d.Name
}
