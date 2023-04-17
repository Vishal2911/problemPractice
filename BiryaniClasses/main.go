package main

import "fmt"


func main() {
	t := 0
	fmt.Print("Enter number of test cases = ")
	fmt.Scanf("%d\n",&t )
	for t >0 {
		x , y := 0,0
		fmt.Print("enter x , y seprated by space = ")
		fmt.Scanf("%d %d\n",&x , &y)
		fmt.Println("Total fee = ", x*y)
		t--
	}
}