package main

import (
	"errors"
	"fmt"
	"net/http"
)

// Interfaces enable us to decouple functionality from implementation
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

// Function types can be called directly
// The LogAdaptor function type meets the Logger interface so we should
// be focusing on making our written functions align to LogAdaptor signature
type LogAdapter func(message string)

type Logic interface {
	SayHello(userID string) (string, error)
}

type SimpleDataStore struct {
	userData map[string]string
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

type Controller struct {
	l     Logger
	logic Logic
}

func main() {

}

// These are business logic function
// Make function name different from interface as PoC
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

// Accept interfaces, return structs
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("calling SayHello for " + userID)

	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}

	return "Hello, " + name, nil
}

// unused but can attach to interface without any problems
func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("calling SayGoodbye for " + userID)

	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}

	return "Goodbye, " + name, nil
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("within SayHello")

	// assume we're storing ids in the query params for ease
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(message))
}

// This just needs to meet the LogAdapter interface
// The LogAdapter interface 'bridges' to the Logger interface
// The fact that LogAdapter is a type means we can use explicit
// type conversions for functions that adopt the same signature
func (lg LogAdapter) Log(message string) {
	lg(message)
}
