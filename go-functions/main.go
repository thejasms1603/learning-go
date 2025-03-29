// package main

// import "fmt"

// func concat(s1 string, s2 string) string {
// 	return s1 + s2
// }

// func main() {
// 	test("Lane,", " happy birthday!")
// 	test("Zuck,", " hope that Metaverse thing works out")
// 	test("Go", " is fantastic")
// }

// func test(s1 string, s2 string){
// 	fmt.Println(concat(s1,s2))
// }

// Passing variables by value
// package main

// func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
// 	lastMonthBill := getBillForMonth(costPerSend, numLastMonth)
// 	thisMonthBill := getBillForMonth(costPerSend, numThisMonth)
// 	return thisMonthBill - lastMonthBill
// }

// func getBillForMonth(costPerSend, messagesSent int) int {
// 	return costPerSend * messagesSent
// }

// Ignoring return values
// we use blank identifier (_)/underscore to ignore the return values
// package main

// func getProductMessage(tier string) string {
// 	quantityMsg, priceMsg, _ := getProductInfo(tier)
// 	return "You get " + quantityMsg + " for " + priceMsg + "."
// }

// func getProductInfo(tier string) (string, string, string) {
// 	if tier == "basic" {
// 		return "1,000 texts per month", "$30 per month", "most popular"
// 	} else if tier == "premium" {
// 		return "50,000 texts per month", "$60 per month", "best value"
// 	} else if tier == "enterprise" {
// 		return "unlimited texts per month", "$100 per month", "customizable"
// 	} else {
// 		return "", "", ""
// 	}
// }

// Named return values
// package main

// func yearsUntilEvents(age int) (int, int, int) {
// 	yearsUntilAdult := 18 - age
// 	if yearsUntilAdult < 0 {
// 		yearsUntilAdult = 0
// 	}
// 	yearsUntilDrinking := 21 - age
// 	if yearsUntilDrinking < 0 {
// 		yearsUntilDrinking = 0
// 	}
// 	yearsUntilCarRental := 25 - age
// 	if yearsUntilCarRental < 0 {
// 		yearsUntilCarRental = 0
// 	}
// 	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
// }

// Naked return values
package main

import "errors"

// initialising return  values with zero values
func getCoords() (x int, y int){
	return // this automatically returns x and y
}

func main() {
	getCoords()
}

func Calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("Can't divide by zero")
	}
	mul = a * b
	div = a / b 
	return mul, div, nil
}

// Explicit Returns
// Even though a function has named return values, we can still explicitly return values if we want to.
func getCoords() (x, y int){
  return x, y // this is explicit
}

func getCoords() (x, y int){
  return 5, 6 // this is explicit, x and y are NOT returned
}

func getCoords() (x, y int){
  return // implicitly returns x and y
}

// Early Returns
// Go supports the ability to return early from a function. This is a powerful feature that can clean up code, especially when used as guard clauses.

// Guard Clauses leverage the ability to return early from a function (or continue through a loop) to make nested conditionals one-dimensional. Instead of using if/else chains, we just return early from the function at the end of each conditional block.

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return dividend/divisor, nil
}