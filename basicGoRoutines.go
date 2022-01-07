package main

import (
	"fmt"
	"time"
)

func myPrint(start, finish int ) {
	for i := start; i <= finish; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	time.Sleep(100* time.Microsecond)
}

func main() {
	for i := 0; i< 5; i++ {
		//A goroutine is the smallest executable Go entity.
		go myPrint(i, 5) // run in a go routine
	}
	time.Sleep(time.Second)
	//A channel in Go is a mechanism that, among other things, allows goroutines to
	// communicate and exchange data.
	// Goroutines and channels, as well as pipelines and sharing data among goroutines
}
