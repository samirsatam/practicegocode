package main

/*
Slices are dynamic, which means that they can grow or shrink after creation if needed.
Any changes you make to a slice inside a function also affect the original slice.

type SliceHeader struct {
	Data uintptr
	Len int
	Cap int
}

*/
import "fmt"

func main() {
	aSlice := []float64{}
	// cap() function - capacity of the slice, cap doubles each time length reaches the cap size.
	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)
	fmt.Println(aSlice, "with length", len(aSlice))

	// A slice with length 4
	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4

	t = append(t, -5)
	fmt.Println(t)

	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}

	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	make2D := make([][]int, 2)
	fmt.Println(make2D)
	make2D[0] = []int{1, 2, 3, 4}
	make2D[1] = []int{-1, -2, -3, -4}
	fmt.Println(make2D)

	// defina anew slice, or you can use make
	old := []int{1, 2, 3, 4}
	// make a new slice with the old slice and some new elements.
	// Note the usage of '...' the spread operator.
	new := append(old, []int{5, 6, 7, 8}...)
	fmt.Println(new)
	
}
