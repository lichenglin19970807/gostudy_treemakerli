package main

import "fmt"

func main() {
	set := make(map[int]int)
	s := []int{1, 2, 3, 1}
	for i, v := range s {
		if e, ok := set[v]; !ok {
			set[v] = i
		} else {
			fmt.Println(e, i)
		}
	}
	fmt.Println(set)
}
