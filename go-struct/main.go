package main

import "fmt"

type messageToSend struct {
	phoneNumber string
	message string
}

func main() {
	test(messageToSend{
		phoneNumber:"123456789",
		message:"Hello, World!",
	})
	test(messageToSend{
		phoneNumber:"987654321",
		message:"Goodbye, World!",
	})
	test (messageToSend{
		phoneNumber: "555555555",
		message: "Hello, Go!",
	})
}

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("===========================")
}

package main

import "fmt"

type car struct {
	Make string
	Model string
	Height int
	Width int
	Color string
	FrontWheel Wheel
	BackWheel Wheel
}

type Wheel struct {
	Radius int
	Material string
}
func main() {
	myCar := car{}
	myCar.Color = "Red"
	myCar.FrontWheel.Radius = 5
	fmt.Printf("My Car Color is: %s\n", myCar.Color)
}

package main

import "fmt"

type messageToSend struct {
	message string
	sender user
	reciever user
}

type user struct {
	name string
	number string
}

func canSendMessage(mToSend messageToSend) bool {
	if(mToSend.sender.name == "" || mToSend.reciever.name== "") {
		return false
	}
	if(mToSend.sender.number == "" || mToSend.reciever.number == ""){
		return false
	}
	return true
}

func main(){
	fmt.Println("Sending message to user")
	fmt.Println(canSendMessage(messageToSend{
		message:  "Hello Bob",
		sender: user{
			name: "Alice",
			number: "123456789",
		},
		reciever: user{
			name:"Bob",
			number:"",
		},
	}))
}

// Anonymous Structs
myCar := struct {
	Make string
	Model string
} {
	Make: "tesla"
	Model: "model 3"
}

// Nested Anonymous Structs
type car struct {
	Make string
	Model string
	Height string
	Width string
	Wheel struct {
		Radius int
		Material string
	}
}

package main

import "fmt"

type rect struct {
	width int32
	height int32
}

func (r rect) area() int32{
	return r.width * r.height
}

func main() {
	r := rect{width : 10, height: 5}
	fmt.Println("Area of rectangle is: ", r.area())
	
}