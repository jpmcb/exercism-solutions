//Package raindrops implements the number to "rainspeak" string
package raindrops

import "strconv"

// Convert takes an int and returns the "rainspeak" string
func Convert(num int) string {
	var rainSpeakString string

	if num%3 == 0 {
		rainSpeakString += "Pling"
	}

	if num%5 == 0 {
		rainSpeakString += "Plang"
	}

	if num%7 == 0 {
		rainSpeakString += "Plong"
	}

	if rainSpeakString == "" {
		return strconv.Itoa(num)
	}
	return rainSpeakString
}
