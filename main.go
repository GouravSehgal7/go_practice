package main

import "fmt"

func main() {
	var value int = 19
	var name = "4"
	num := 10
	fmt.Println(value)
	fmt.Printf("%T", value)
	fmt.Println(name)
	fmt.Printf("%T", name)
	fmt.Println(num)
	fmt.Printf("%T", num)
	var array = [3]int{}
	array[0] = 3
	array[1] = 4
	fmt.Println(array)
	var arr = []int{}
	fmt.Printf("%T", arr)
	arr = append(arr, 2, 3, 4, 5, 6)
	fmt.Println(arr)

}
