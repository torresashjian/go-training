package embedding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCastit(t *testing.T) {
	// create new Inside
	inside := MyInsideType{Name: "Inside Name"}
	//outside := MyOutsideType{inside: &inside}
	name, err := Castit(&inside)
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(t, "Inside Name", name)
}
