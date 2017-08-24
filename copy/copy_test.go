package copy

import "testing"

func TestCopy(t *testing.T) {
	dummy := &Dummy{name: "dummyName", age: 6}

	t.Logf("Dummy Before change: '%+v'", dummy)

	// Make copy
	cDummy1, cDummy2 := dummy.copy()

	cDummy1.age = 0
	cDummy2.age = 1

	t.Logf("Dummy After change: '%+v'", dummy)
	t.Logf("Dummy1 Copy After change: '%+v'", cDummy1)
	t.Logf("Dummy2 Copy After change: '%+v'", cDummy2)

}
