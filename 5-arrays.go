package main

import "fmt"

func main() {
	// Defining fixed size array
	ar := []int{1, 2, 3, 4}
	fmt.Println(ar)
	fmt.Println("Size of array: ", len(ar))

	// Defining a variable length array (or slice), similar to ArrayList
	sl := make([]int, 4)
	fmt.Println(sl)
	fmt.Println("Size of slice: ", len(sl))
	// NOTE: append need assignment
	sl = append(sl, 1)
	fmt.Println(sl)
	fmt.Println("Size of slice after append: ", len(sl))

	// Iterate over array|slice
	for idx, val := range sl {
		fmt.Println("Idx: ", idx, "Val: ", val)
	}

	// Let's see 2D array now
	var ar2d [2][2]float32
	ar2d[0][0] = 0.0
	ar2d[0][1] = 0.1
	ar2d[1][0] = 1.0
	ar2d[1][1] = 1.1

	fmt.Println(ar2d)

	// How about 2d slices
	sl2d := make([][]int, 3)
	const sizeArr = 4
	for i := 0; i < len(sl2d); i++ {
		sl2d[i] = make([]int, sizeArr)
		for j := 0; j < sizeArr; j++ {
			sl2d[i][j] = i * sizeArr + j
		}
	}
	fmt.Println(sl2d)
}
