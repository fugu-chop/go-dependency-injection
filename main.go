package main

import "fmt"

// Interfaces enable us to decouple functionality from implementation
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type LogAdapter func(message string)

type SimpleDataStore struct {
	userData map[string]string
}

func main() {

}

func LogOutput(message string) {
	fmt.Println(message)
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// lg is a function type, so we can call it
func (lg LogAdapter) Log(message string) {
	lg(message)
}
