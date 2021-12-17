package main

import (
	"fmt"
)

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	Ints := []int{1,2,3}
	Strings := []string{"One", "Two", "Three"}
	Print(Ints)
	Print(Strings)
}

show index from non_mdm_event_store;
ALTER TABLE non_mdm_event_store DROP INDEX non_mdm_event_store_id_index;
