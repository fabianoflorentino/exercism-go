package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	welcome_message := fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))

	return welcome_message
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	add_border := fmt.Sprintf("%s\n%s\n%s", strings.Repeat("*", numStarsPerLine), welcomeMsg, strings.Repeat("*", numStarsPerLine))

	return add_border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanup_message := strings.ReplaceAll(oldMsg, "*", "")

	return strings.TrimSpace(cleanup_message)
}
