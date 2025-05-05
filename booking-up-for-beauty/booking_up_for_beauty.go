package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	parsedTime, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic("Invalid date format")
	}

	return parsedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	appointmentPassed := false
	currentTime := time.Now()

	parsedTime, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic("Invalid date format")
	}

	if parsedTime.Before(currentTime) {
		appointmentPassed = true
	}

	return appointmentPassed
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsedTime, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic("Invalid date format")
	}

	return parsedTime.Hour() >= 12 && parsedTime.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsedTime, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic("Invalid date format")
	}

	appointmentDescription := fmt.Sprintf("You have an appointment on %s, %s, at %02d:%02d.", parsedTime.Weekday(), parsedTime.Format("January 2, 2006"), parsedTime.Hour(), parsedTime.Minute())

	return appointmentDescription
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
