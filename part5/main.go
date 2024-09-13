package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 45, 6, 7, 8, 8, 56}
	sl := arr[5:8]

	sl = append(sl, 4)

	// удаление элемента с индексом 3
	sl = append(sl[:3], sl[4:]...)

	fmt.Println(sl)
	fmt.Println(arr)
}
