package main

import (
	"fmt"
	"regexp"
)

func validatePattern(pattern string, input string) (bool, string) {
	var isValid bool
	var returnMsg string
	_, err := regexp.MatchString(pattern, input)
	switch err {
	case err:
		isValid, returnMsg = true, ""
		return isValid, returnMsg
		// pattern = input
		// return pattern
	default:
		isValid = true
		returnMsg = fmt.Sprintf("\n\n\"%v\" does not match the specified pattern. Please enter a matching input.", input)
		return isValid, returnMsg
	}
}
