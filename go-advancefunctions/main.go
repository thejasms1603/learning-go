// Functions as arguments
// package main

// import "fmt"

// func sum(x int, y int) int {
// 	return x + y
// }

// func sub(a int, b int) int{
// 	return a - b
// }

// func calculate(x int, y int, fn func(int, int) int) int {
// 	return fn(x,y)
// }
// func main(){
// 	sum := calculate(4,5,sum)
// 	sub := calculate(3,1,sub)
// 	fmt.Printf("Sum is %v and Sub is %v", sum, sub)
// }

//Returning Functions

package main

import "fmt" 

func multiplier(factor int) func(int) int {
	return func(a int) int {
		return a*factor
	}
}

func main() {
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(3))
	fmt.Println(triple(3))

}