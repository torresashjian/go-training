package flags

import (
	"flag"
	"encoding/json"
)

type MyStruct struct {
	Name string
}
func setStructFlag() string{
	//h := json.RawMessage(`{"precomputed": true}`)
	//raw := json.RawMessage(`{"foo":"bar"}`)
	var nilRaw []byte
	j, err := json.Marshal(&nilRaw)
	if err != nil {
		panic(err)
	}
	flag.String("myjsonRaw", string(j), "my json raw")
	flag.Parse()

	f := flag.Lookup("myjsonRaw")
	return f.Value.String()
}