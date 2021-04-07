package main

import "fmt"

func transpose(a [5][5]int) [5][5]int {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a[i]); j++ {
			x := a[i][j]
			a[i][j] = a[j][i]
			a[j][i] = x
		}
	}
	return a
}

func rotate(a [5][5]int) [5][5]int {
	for i := 0; i < len(a); i++ {
		left := 0
		right := len(a[i]) - 1
		for left < right {
			x := a[i][left]
			a[i][left] = a[i][right]
			a[i][right] = x
			left++
			right--
		}
	}

	return a
}

func main() {
	a := [5][5]int{}

	k := 0
	for i := range a {
		for j := range a[i] {
			a[i][j] = k
			k += 1
		}
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Print(a[i][j])
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}

	a = transpose(a)
	a = rotate(a)
	fmt.Println("========ROATATED============")
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Print(a[i][j])
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}

}
