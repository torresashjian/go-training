package types

import "fmt"

type MyType int
func(t *MyType) Method1(){
	fmt.Println("Test here")
}

type AnotherType MyType

type PtrAnotherType *MyType

type EmbAnotherType struct {
	MyType
}


func TestTypes(){
	m := new(MyType)
	m.Method1()

	//at := new(AnotherType)
	//at.Method1() //INVALID

	//t := new(PtrAnotherType)
	//t.Method1() //INVALID

	et := new(EmbAnotherType)
	et.Method1()
}
