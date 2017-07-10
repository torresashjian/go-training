package copy


type Dummy struct {
	name string
	age int
}

func (d *Dummy) copy () (*Dummy, *Dummy){
	cDummy1 := *d
	cDummy2 := *d
	return &cDummy1, &cDummy2
}
