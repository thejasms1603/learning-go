package main

import (
	"errors"
	"fmt"
)

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	u, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if u.scheduledForDeletetion {
		delete(users,name)
		return true,nil
	}
	return false, nil
}

type user struct {
	name string
	number int
	scheduledForDeletetion bool
}

func getCounts(userIDs []string) map[string]int {
	counts := make(map[string]int)
	for _, userID := range userIDs {
		count := counts[userID]
		count++
		counts[userID] = count
	}
	return counts
}

func main() {
	m:= map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}

	delete(m, "two")

	value, exists := m["two"]
	if exists {
		fmt.Println("Value found:", value)	
	} else {
		fmt.Println("Value not found")
	}

	m["four"] = 4

	for k,v := range m {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

}