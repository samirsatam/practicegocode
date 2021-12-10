package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	fmt.Println(os.TempDir())
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// A call to OpenFile creates a log file for writing, if the file doe not exist.
	// else OpenFile opens it for writing/reading. O_APPEND appends new data.

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	// BASIC
	// iLog := log.New(f, "iLog ", log.LstdFlags)
	// iLog.Println("Hello There!")

	//Date filename/linenumber
	LstdFlags := log.Ldate | log.Lshortfile
	iLog := log.New(f, "LNum ", LstdFlags)
	iLog.Println("Mastering Go 3rd edition!")

	//Date Time filename/linenumber
	iLog.SetFlags(log.Lshortfile | log.LstdFlags)
	iLog.Println("Another log entry!")
}
