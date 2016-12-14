package subpackage

import (
	"go/build"
	"reflect"
)

type my_type interface {
	Name() string
}

// This method should just using the object passed
// find the package.json file that lives in the same dir
func Register(t my_type) string {
	t_type := reflect.TypeOf(t)
	t_path := t_type.PkgPath()
	ctx := build.Default
	pkg, err := ctx.Import(t_path, ".", build.FindOnly)
	if err != nil {
		return err.Error()
	}
	return pkg.Dir
}
