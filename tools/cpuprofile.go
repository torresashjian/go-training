package tools

import (
	//"flag"
	"fmt"
	"time"
	"net/http"
	"log"

	//"github.com/pkg/profile"
	_ "net/http/pprof"
)

//var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func Start() {
	fmt.Println("Inside Start()")
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	//defer profile.Start(profile.CPUProfile).Stop()
	fib := Fib(10)
	fmt.Println("Fib %d", fib)
}

var alreadyStopped bool = false

func Fib(n int) int {
	if n < 2 {
		return n
	}
	if n == 10 && !alreadyStopped {
		alreadyStopped = true
		fmt.Println("Stopping")
		time.Sleep(60 * time.Second)
		fmt.Println("End Stopping")
	}
	return Fib(n-1) + Fib(n-2)
}
