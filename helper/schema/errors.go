package schema

import (
	"fmt"
)

type AttributeError struct {
	attributeAddr string // dotted syntax
	msg           string
}

func (ae *AttributeError) Error() string {
	return fmt.Sprintf("%s: %s", ae.attributeAddr, ae.msg)
}
