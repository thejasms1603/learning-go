
// In Javascript, we use try and catch block  and the function goes like this
// function main () {
// 	try {
// 		const user = getUser()
// 	} catch(err)
// 	{
// 		console.log(err)
// 	}
// }

// error handling in go
package main 

import "fmt"
func main() {
	user, err := getUser()
	if err != nil {
		fmt.Println(user)
	}
	profile, err := getUserProfile(user.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getUser() (User, error) {
// do some get user logic here

}

package main

import "fmt"

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	costForCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	costForSpouse, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	return costForCustomer + costForSpouse, nil

}

func sendSMS(message string) (float64, error) {
	const maxTextLen = 32
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("Cant send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}

The Error Interface
type error Interface {
	Error() string
}

// Formatting strings review
// A convinient way to format strings in Go is by using standard's library fmt.sprintf function()

package main

import "fmt"

func main() {
	name := "Thejas"
	age := 22
	s:= fmt.Sprintf("My name is %v, and I am %v years old", name, age)
	fmt.Println(s)
}

// The %v verb is used to format the value in a default format. It is a versatile verb that can be used with any data type.
// The %T verb is used to format the type of the value. It is useful for debugging and understanding the type of a variable.
// The %d verb is used to format integers. It is a common verb used for formatting numeric values.
// The %f verb is used to format floating-point numbers. It is useful for displaying decimal values.
// The %s verb is used to format strings. It is the most common verb used for formatting string values.

package main

import "fmt"

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs %.2f to be sent to '%v' can not be sent", cost, recipient)
}

func test(cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
	fmt.Println("==================")
}

func main (){
	test(1.4, "+1 (435) 555 0923")
	test(2.1, "+2 (702) 555 3452")
	test(1.4, "+1 (917) 671 8140")

}

// Building Custom Error Interface

package main

import "fmt"
type userError struct {
	name string
}

func (e userError) Error() string {
	return fmt.Sprintf("%v has a problem with their account", e.name)
}

func sendSMS(msg, userName string) error {
	if !canSendToUser(userName) {
		return userError{name: userName}
	}
}

package main

import "fmt"

type divideError struct {
	dividend float64
}

func (e divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero",e.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}
func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

// Errors Package
package main
import "errors"

var err error = errors.New("something went wrong")