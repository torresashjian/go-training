package embedding

import (
	"errors"
)

type MyInsideType struct {
	Name string
}

type MyOutsideType struct {
	inside *MyInsideType
}

type Context interface {
	GetName() string
}

func (it MyInsideType) GetName() string {
	return it.Name
}

func (ot MyOutsideType) GetName() string {
	return ot.inside.Name
}

func Castit(ctx Context) (string, error) {
	// Fist cast to MyOutsideType
	if myOutside, ok := ctx.(*MyOutsideType); ok {
		return myOutside.inside.Name, nil
	}
	return "", errors.New("No cast possible")
}
