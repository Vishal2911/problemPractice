package main

import "fmt"


func main() {

t := 0

fmt.Print("Enter number of test cases = ")
fmt.Scanf("%d\n",&t)

for t > 0 {
	x := 0 
	fmt.Print("Enter frequency of command = ")
	fmt.Scanf("%d\n",&x)

	if x >= 67 && x <= 45000 {
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
	t--
}

}