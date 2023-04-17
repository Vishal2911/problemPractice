package main

import "fmt"

func main() {
	t := 0
	fmt.Print("enter number of test cases = ")
	fmt.Scanf("%d\n", &t)

	// for i:= 1 ; i <= t ; i++ {
	// 	x , y , a := 0,0,0
	// 	fmt.Print("enter min age = ")
	// 	fmt.Scanf("%d\n",&x)
	// 	fmt.Print("enter max age = ")
	// 	fmt.Scanf("%d\n",&y)
	// 	fmt.Print("enter user age =")
	// 	fmt.Scanf("%d\n",&a)
	// 	if a >= x && a < y {
	// 		fmt.Println("YES")
	// 	}else{
	// 		fmt.Println("NO")
	// 	}
	// }

	for t > 0 {
		x, y, a := 0, 0, 0
		fmt.Print("enter min age = ")
		fmt.Scanf("%d\n", &x)
		fmt.Print("enter max age = ")
		fmt.Scanf("%d\n", &y)
		fmt.Print("enter user age =")
		fmt.Scanf("%d\n", &a)
		if a >= x && a < y {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		t--
	}

}
