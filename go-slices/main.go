// Empty Slice
package main

func main() {
	var users []string
	println(users == nil)
}

// Initialising slice
package main

import "fmt"

func main() {
	var users []string = []string{"Thejas","Bhuvan"}
	fmt.Print(users)
}

// Initialize without default value (make)
package main

import "fmt"

func main() {
	var users []string = make([]string,3)
	println(len(users))
	print(cap(users))
	fmt.Print(users[0]=="")
}

// Copy by reference
package main

import "fmt"

func main() {
	var users []string = []string{"harkirat", "raman"}
	var users2 = users // Copied by reference
	users2[0] = "harkirat2"
	fmt.Print(users)
}

// Copy by value
package main

import "fmt"

func main() {
	var users []string = []string{"harkirat", "raman"}
	var users2 = make([]string, len(users))

	copy(users2, users)
	users2[0] = "harkirat2"
	fmt.Print(users)
}


package main

import (
	"errors"
	"fmt"
)

const (
	planPro  = "pro"
	planFree = "free"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	if plan == planPro {
		return allMessages[:], nil
	}
	if plan == planFree {
		return allMessages[0:2], nil
	}
	return nil, errors.New("unsupported plan")
}

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func test(name string, doneAt int, plan string) {
	defer fmt.Println("================")
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages, err := getMessageWithRetriesForPlan(plan)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf("sending: %v", msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("No response")
		}
	}
}

func main() {
	test("Ozgur", 3, planFree)
	test("Jeff", 3, planPro)
	test("Sally", 2, planPro)
	test("Sally", 3, "no plan")
}


package main

import ( "fmt" )

func main() {
	var users []string
	if( users == nil) {
		fmt.Println("users is nil")
	} else {
		fmt.Println("users is not nil")
	}
}