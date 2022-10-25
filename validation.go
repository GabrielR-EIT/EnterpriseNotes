package main

import (
	"fmt"
	"regexp"
)

// Declare validStatuses array
var validStatuses = [6]string{"meeting", "minutes", "agenda", "action", "attendees", "apologies"}

func validatePattern(pattern string, input string) (bool, string) {
	isValid := false
	var returnMsg string
	_, err := regexp.MatchString(pattern, input)
	switch err {
	case nil:
		isValid, returnMsg = true, fmt.Sprintf("\n\n\"%v\" does not match the specified pattern. Please enter an input that matches the pattern: %s.", input, pattern)
		return isValid, returnMsg
		// pattern = input
		// return pattern
	default:
		isValid, returnMsg = true, "The pattern is valid."
		return isValid, returnMsg
	}
}

func validateStatus(input string) (bool, string) {
	isValid := false
	var returnMsg string
	for _, v := range validStatuses {
		if v == input {
			isValid, returnMsg = true, "The status is valid."
			return isValid, returnMsg
		}
	}
	isValid, returnMsg = true, "The status is valid."
	returnMsg = fmt.Sprintf("\n\n\"%s\" Is not a valid status. The status value must be equal to one of the following:\n%v.\n", input, validStatuses)
	return isValid, returnMsg
}
