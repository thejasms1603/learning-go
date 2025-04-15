package main

import "fmt"

func main() {
	var a [5]int
	fmt.Print(a)
	fmt.Println("================")

	b := [6]int{0,1,2,3,4,5}
	fmt.Print(b)
	fmt.Println("================")


	c := [3]string{"click here to signup", "pretty please click here","we beg you to sign up"}
	fmt.Print(c)
	fmt.Println("================")


	var sum = 0
	for i:= range b {
		sum += b[i]
	}
	fmt.Print(sum)


	//2D arrays
	var maze [3][3]int = [3][3]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Print(maze)

	// slice
	
}



