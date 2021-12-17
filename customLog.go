package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"reflect"
)

func mainOld() {
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

func arbitArgs(args ...interface{}) {
	var total, nInts, nFloats, nStrings int
	invalid := make([]string, 0)
	// fmt.Println(args...) //Print all the args
	// fmt.Println(reflect.TypeOf(args).Kind()) // is a slice
	for _, k := range args {
		// fmt.Println("------------------")
		// fmt.Println(reflect.TypeOf(k).Kind()) // is slice inside the slice above
		// fmt.Println(k)
		// fmt.Println("------------------")
		// strconv.Atoi(k.(string)) is used to convert string to ints,
		// strconv.ParseFloat(k.(string), 64) is used to convert string to floats,
		// this is how it was originally supposed tp work, but we somehow stumbled on Reflection.
		switch reflect.TypeOf(k).Kind() {
		case reflect.Slice:
			{
				s := reflect.ValueOf(k) // need to ValueOf to make a interface {} to a slice
				for i := 0; i < s.Len(); i++ {
					// fmt.Println(reflect.ValueOf(s.Index(i)).Kind()) // returns a struct, not sure about this.
					// fmt.Println(s.Index(i).Kind())                  // identifies correctly as string.
					var argType = s.Index(i).Kind()
					switch argType {
					case reflect.String:
						{
								total++
								nStrings++
								continue
						}
					case reflect.Int:
						{
							total++
							nInts++
							continue

						}
					case reflect.Float32:
						{
								total++
								nFloats++
								continue
						}
					}
				}
			}

		}
	}

	fmt.Println("------------------------")
	fmt.Println(total)
	fmt.Println(nInts)
	fmt.Println(nFloats)
	fmt.Println(nStrings)
	fmt.Println(invalid)
}

func main() {
	// errorVariables(1, "2", 3)
	arbitArgs(os.Args)
}
