package blackjack

import (
	"fmt"
	"os"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	c, err := cardValidation(card)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	}

	return c
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	f, err := decisionFirstTurn(card1, card2, dealerCard)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	}

	return f
}

// decisionFirstTurn determines the player's decision for the first turn in a game of blackjack
// based on the values of the player's two cards and the dealer's card.
//
// Parameters:
//   - card1: A string representing the player's first card.
//   - card2: A string representing the player's second card.
//   - dealerCard: A string representing the dealer's card.
//
// Returns:
//   - A string representing the player's decision:
//     "P" for "Split", "H" for "Hit", "S" for "Stand", or "W" for "Win".
//   - An error if any of the card values are invalid.
func decisionFirstTurn(card1, card2, dealerCard string) (string, error) {
	c1, err := cardValidation(card1)
	if err != nil {
		return "", fmt.Errorf("invalid card1: %v", err)
	}

	c2, err := cardValidation(card2)
	if err != nil {
		return "", fmt.Errorf("invalid card2: %v", err)
	}

	d, err := cardValidation(dealerCard)
	if err != nil {
		return "", fmt.Errorf("invalid dealerCard: %v", err)
	}

	if handScore(c1, c2) == 22 {
		return "P", nil
	}

	if handScore(c1, c2) == 21 {
		if d < 10 {
			return "W", nil
		} else {
			return "S", nil
		}
	}

	if handScore(c1, c2) >= 17 || (handScore(c1, c2) >= 12 && d < 7) {
		return "S", nil
	}

	return "H", nil
}

// deckValidation takes a card name as a string and returns its corresponding
// value in a standard deck of cards. The card values are defined as follows:
// "ace" is 11, "two" through "ten" are their respective numeric values,
// and "jack", "queen", and "king" are 10. If the card is not found in the
// predefined deck, an error is returned.
//
// Parameters:
//   - card: A string representing the name of the card.
//
// Returns:
//   - An integer representing the value of the card.
//   - An error if the card is not found in the deck.
func cardValidation(card string) (int, error) {
	c := map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
	}

	return c[card], nil
}

// handScore calculates the score of a hand of two cards.
func handScore(card1, card2 int) int {
	return card1 + card2
}
