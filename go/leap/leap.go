// Package leap exposes a method to check if a given year is a leap year
package leap

// IsLeapYear returns true if the given int is a leap year
func IsLeapYear(year int) bool {

	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}

	return false
}
