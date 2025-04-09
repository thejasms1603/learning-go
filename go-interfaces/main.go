// package main

// import "math"

// type shape interface {
// 	area() float64
// 	perimeter() float64
// }

// type rect struct {
// 	width, height float64
// }

// func (r rect) area() float64 {
// 	return r.width * r.height
// }

// func (r rect) perimeter() float64 {
// 	return 2*r.width * 2*r.height
// }

// type circle struct {
// 	radius float64
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func (c circle) perimeter() float64 {
// 	return 2* math.Pi * c.radius
// }

// func main() {

// }

package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) (string, int) {
	text := msg.getMessage()
	fmt.Println(text)
	return text, len(text) * 2
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func main() {
	bm := birthdayMessage{
		birthdayTime:  time.Now().AddDate(0, 0, 7),
		recipientName: "Alice",
	}

	sr := sendingReport{
		reportName:    "Weekly Stats",
		numberOfSends: 42,
	}
	sendMessage(bm)
	sendMessage(sr)
}