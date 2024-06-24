package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	welcome := fmt.Sprintf("Welcome to my party, %s!", name)

	return welcome
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	happy_birdthday := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)

	return happy_birdthday
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	assign_table := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
	directions := fmt.Sprintf("You will be sitting next to %s.", neighbor)

	return Welcome(name) + "\n" + assign_table + "\n" + directions
}
