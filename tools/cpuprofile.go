package tools

import (
	//"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	//"github.com/pkg/profile"
	"io/ioutil"
	_ "net/http/pprof"
	"os"
	"strconv"
)

//var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func Start() {
	fmt.Println("Inside Start()")
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	//defer profile.Start(profile.CPUProfile).Stop()
	fib := Fib(60)
	fmt.Println("Fib %d", fib)
}

var alreadyStopped bool = false

func Fib(n int) int {
	if n < 2 {
		return n
	}
	writeToFile(n)
	return Fib(n-1) + Fib(n-2)
}

// I know this method is way too stupid, just adding some random computation
func writeToFile(n int) {
	fmt.Println("Writting to file")
	time.Sleep(1 * time.Second)
	counter := 0
	for i := 0; i < n; i++ {
		counter = i
	}
	content := []byte(strconv.Itoa(counter))
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
