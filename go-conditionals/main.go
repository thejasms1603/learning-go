// IF and IF Else

package main

func main(){
	a := 2
	if a % 2 == 0 {
		println("A is even")
	} else {
		println("A is odd")
	}
}

package main

import "fmt"

func main() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	finalCost = bulkMessageCost
	if isPremiumUser {
		finalCost = bulkMessageCost - (bulkMessageCost * discountRate)
	} else {
		finalCost = bulkMessageCost
	}

	if finalCost <= accountBalance {
		accountBalance = accountBalance - finalCost
		fmt.Println(purchaseSuccessMessage)
	} else {
		fmt.Println(insufficientFundMessage)
	}

	fmt.Println("Account balance:", accountBalance)
}

// FOR LOOP
package main

func main() {
	sum := 0
	for i:=1; i < 10;i++ {
		sum += i
	}
	println(sum)
}


func main() {
	sum := 0
	for i :=  range 10 {
		sum += i
	}
	println(sum)
}



// SWITCH CASE
package main

import "fmt"

func billingCost(plan string) float64 {
	switch plan {
	case "basic":
		return 10.0
	case "pro":
		return 20.0
	case "enterprise":
		return 50.0
	default:
		return 0.0
	}
}


func main() {
	plan := "basic"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "pro"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "enterprise"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "free"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "unknown"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
}