// Get a file in same package
package files

import "github.com/torresashjian/go-training/files/subpackage"

type my_type int

func (t my_type) Name() string {
	return "my name"
}

func registration() string {
	var t my_type = 1
	return subpackage.Register(t)
}
