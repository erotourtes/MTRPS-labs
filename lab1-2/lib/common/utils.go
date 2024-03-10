package common

import "regexp"

func RemoveRepeatedSpaces(s string) string {
	re := regexp.MustCompile(`\s+`)
	output := re.ReplaceAllString(s, " ")
	return output
}
