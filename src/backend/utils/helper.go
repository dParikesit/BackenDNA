package utils

import "strings"

func SplitText(text string, delimiter string) []string {
	return strings.Split(text, delimiter)
}

func JoinArray(arr []string, idx int) string {
	var name string = ""
	for i := idx; i < len(arr); i++ {
		name += arr[i] + " "
	}
	return strings.TrimSpace(name)
}
