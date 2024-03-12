package logwrapper

import (
	uuid "github.com/google/uuid"
)

type Span struct {
	ID     string
	parent *Span
}

func createSpan(parent *Span) *Span {
	s := &Span{
		ID:     uuid.New().String(),
		parent: parent,
	}
	return s
}
