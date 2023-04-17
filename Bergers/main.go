package main

import "fmt"

func main() {
	t := 0
	fmt.Print("Enter number of testcases=")
	fmt.Scanf("%d", &t)
	for t > 0 {
		a, b := 0, 0
		fmt.Print("enter number of patties = ")
		fmt.Scanf("%d\n", &a)
		fmt.Print("Enter number of buns = ")
		fmt.Scanf("%d\n", &b)

		if a < b {
			fmt.Println(a)
		} else {
			fmt.Println(b)
		}
		t--
	}
}
