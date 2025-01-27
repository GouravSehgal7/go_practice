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
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	str := "hello"
	for index, val := range str {
		fmt.Printf("Index: %d, Char: %c\n", index, val)
	}

	n := map[string]int{"Alice": 30, "Bob": 25}
	fmt.Println(n)

	dict := map[string]int{}
	dict["A"] = 1
	dict["B"] = 2
	for i, j := range dict {
		fmt.Println(i, "-->", j)
	}

	k := map[string]int{"Alice": 30, "Bob": 25}
	for key, value := range k {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	m := make(map[U]P)
	user1 := U{name: "name1", email: "jsoie", phone: 12432}
	parm1 := P{haspower: true, isauth: true}

	user2 := U{name: "nagyme1", email: "gtj", phone: 8765}
	parm2 := P{haspower: true, isauth: false}

	m[user1] = parm1
	m[user2] = parm2

	fmt.Println(m)

	for k, l := range m {
		fmt.Println(k, "==>>", l)
	}

	for user, perm := range m {
		fmt.Printf("User: %s, Email: %s, Phone: %d, Has Power: %t, Is Auth: %t\n",
			user.name, user.email, user.phone, perm.haspower, perm.isauth)
	}

	if perm, exists := m[user1]; exists {
		fmt.Printf("\nPermissions for %s: %+v\n", user1.name, perm)
	}

}

type U struct {
	name  string
	email string
	phone int
}

type P struct {
	haspower bool
	isauth   bool
}
