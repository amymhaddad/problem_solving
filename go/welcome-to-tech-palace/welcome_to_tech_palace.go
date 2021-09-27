package techpalace

import (
	"fmt"
	"regexp"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {

	firstHalf, secondHalf := numStarsPerLine, numStarsPerLine

	var greeting string
	for firstHalf >= 0 {
		greeting += "*"
		firstHalf--
	}

	greeting += fmt.Sprintf("\n%s\n", welcomeMsg)

	for secondHalf >= 0 {
		greeting += "*"
		secondHalf--
	}
	fmt.Println(greeting)
	return greeting
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	re := regexp.MustCompile(`([A-Z]+\s?.?)`)
	result := re.findStringSubmatch(oldMsg)
	fmt.Println(result)
	return "here"
}
