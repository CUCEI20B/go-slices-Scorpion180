package main

import "fmt"

func main() {
	var size int = 0
	var i int = 0
	var sum int = 0
	var acumm int = 0
	var list []int

	fmt.Scanln(&size)

	for i < size {
		fmt.Scanln(&sum)
		list = append(list, sum)
		acumm = acumm + sum
		i = i + 1
	}

	fmt.Print(acumm)

}
