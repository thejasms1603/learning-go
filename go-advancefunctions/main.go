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

// package main

// import "fmt"

// func multiplier(factor int) func(int) int {
// 	return func(a int) int {
// 		return a*factor
// 	}
// }

// func main() {
// 	double := multiplier(2)
// 	triple := multiplier(3)

// 	fmt.Println(double(3))
// 	fmt.Println(triple(3))

// }

// Defer

// package main

// type user struct {
// 	name  string
// 	admin bool
// }

// const (
// 	logDeleted  = "User deleted"
// 	logNotFound = "User not found"
// 	logAdmin    = "User is an admin"
// )

// func logAndDelete(users map[string]user, name string) (log string) {
// 	defer delete(users, name)
// 	user, ok := users[name]
// 	if !ok {
// 		return logNotFound
// 	}
// 	if user.admin {
// 		return logAdmin
// 	}
// 	return logDeleted
// }

//Closures in Go

// package main

// import "fmt"

// func concatter() func(string) string {
// 	doc := ""
// 	return func(word string) string {
// 		doc += word + " "
// 		return doc
// 	}
// }

// func outer() func(int) int{
// 	sum:= 0
// 	return func(x int) int{
// 		sum+= x
// 		return sum
// 	}
// }

// func main() {
// 	harryPotterAggregator := concatter()
// 	harryPotterAggregator("Harry")
// 	harryPotterAggregator("Potter")
// 	harryPotterAggregator("and")
// 	harryPotterAggregator("the")
// 	harryPotterAggregator("Goblet")
// 	harryPotterAggregator("of")
// 	harryPotterAggregator("Fire")
// 	harryPotterAggregator("and")
// 	harryPotterAggregator("the")
// 	harryPotterAggregator("Order")
// 	harryPotterAggregator("of")
// 	harryPotterAggregator("the")
// 	fmt.Println(harryPotterAggregator("Phoenix"))

// 	sum := outer()
// 	sum(1)
// 	sum(2)
// 	sum(3)
// 	sum(4)
// 	fmt.Println(sum(5))
// }

// Anonymous functions

package main

import "fmt"

func main() {
	sum, sub := func(a,b int) (sum int, sub int) {
		sum = a + b
		sub = a - b
		return
	}(2,3)

	fmt.Println(sum , " ", sub)
}

