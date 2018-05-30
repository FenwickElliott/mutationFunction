package main

import "fmt"

func main() {
	i := 0
	fmt.Println(i)
	increment(&i)
	fmt.Println(i)
}

func increment(i *int) {
	*i++
}
