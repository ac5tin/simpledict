package dict

import "errors"

type Result struct {
	Definition []string
	Synonyms   []string
}

type Dict interface {
	GetDefinition(s *string) (*Result, error)
}

type dict struct {
	Dict
}

func (d *dict) GetDefinition(s *string) (*Result, error) {
	return nil, errors.New("not implemented")
}
