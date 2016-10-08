package files

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// This function changes the creation time of the list of files secuentially
func Change(src, dst string) {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		// Read all content of src to data
		data, err := ioutil.ReadFile(filepath.Join(src, file.Name()))
		checkErr(err)
		// Write data to dst
		err = ioutil.WriteFile(filepath.Join(dst, file.Name()), data, 0644)
		checkErr(err)
	}
}
