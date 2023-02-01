package utils

import "regexp"

var (
	md532Pattern = regexp.MustCompile("^[0-9a-fA-F]{32}$")
)

func IsMD5(text string) bool {
	return md532Pattern.MatchString(text)
}
