package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and stands out his age.
func HappyBirthday(name string, age int) string {
	panic("Please implement the HappyBirthday function")
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbour, direction string, distance float64) string {
	panic("Please implement the AssignTable function")
}
