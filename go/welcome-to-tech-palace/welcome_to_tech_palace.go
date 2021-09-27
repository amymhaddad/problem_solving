package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var greeting string
	greeting = strings.Repeat("*", numStarsPerLine)
	greeting += "\n" +welcomeMsg + "\n"
	// greeting = strings.Repeat("\n", 1)
	// greeting = strings.Repeat(welcomeMsg, 1)
	// greeting = strings.Repeat("\n", 1)
//	fmt.Println("H: ", greeting)
	greeting = strings.Repeat("*", numStarsPerLine)
	return greeting

	// firstHalf, secondHalf := numStarsPerLine, numStarsPerLine
	//
	// var greeting string
	// for firstHalf >= 0 {
	// 	greeting += "*"
	// 	firstHalf--
	// }
	//
	// greeting += fmt.Sprintf("\n%s\n", welcomeMsg)
	//
	// for secondHalf >= 0 {
	// 	greeting += "*"
	// 	secondHalf--
	// }
	// fmt.Println(greeting)
	// return greeting
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.ReplaceAll(oldMsg, "*", "")
	return strings.TrimSpace(oldMsg)
}
