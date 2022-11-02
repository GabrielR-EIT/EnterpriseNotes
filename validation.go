package main

import (
	"fmt"
	"log"
	"regexp"
)

// Declare validStatuses array
var validStatuses = [5]string{"none", "in progress", "completed", "cancelled", "delegated"}

// Declare Patterns array
var Patterns = map[string]string{`[a-zA-z]+`: "A sentence with a given prefix and/or suffix",
	`[0-9\W]`: "A phone number with a given area code and optionally a consecutive sequence of numbers that are part of that number",
	`@{1}`:    "An email address on a domain that is only partially provided",
	`meeting|minutes|agenda|action|attendees|apologies{3,}`: "Text that contains at least three of the following case-insensitive words: meeting, minutes, agenda, action, attendees, apologies",
	`[A-Z]{3,}`: "A word in all capitals of three characters or more"}

func validatePattern(pattern string, input string) (bool, string) {
	isValid, returnMsg := false, ""
	isValid, err := regexp.MatchString(pattern, input)
	if err != nil {
		log.Printf("An error occurred when validating the pattern.\nGot %s\n", err)
		return isValid, returnMsg
	}
	if !isValid {
		returnMsg = fmt.Sprintf("\n\n\"%v\" does not match the specified pattern. Please enter an input that matches the pattern: %s.", input, pattern)
		return isValid, returnMsg
	}
	returnMsg = "The pattern is valid."
	return isValid, returnMsg
}

// Note: Pass values as lowercase to validateStatus() function
func validateStatus(input string) (bool, string) {
	isValid, returnMsg := false, ""
	for _, v := range validStatuses {
		if v == input {
			isValid, returnMsg = true, "The status is valid."
			return isValid, returnMsg
		}
	}
	returnMsg = fmt.Sprintf("\n\n\"%s\" Is not a valid status. The status value must be equal to one of the following:\n%v.\n", input, validStatuses)
	return isValid, returnMsg
}
