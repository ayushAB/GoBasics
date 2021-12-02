package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	//when no array length is given then it is a slice
	primes := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(primes)

	primes = append(primes, 17)

	fmt.Println(primes)
}
