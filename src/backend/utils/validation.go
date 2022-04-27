package utils

import (
	"regexp"
)

var dnaRegex *regexp.Regexp = regexp.MustCompile("^[ATCG]+$")
var diseaseRegex *regexp.Regexp = regexp.MustCompile(`^[ A-Za-z0-9-\s]*$`)

func IsValidDNA(s string) bool {
	return dnaRegex.MatchString(s)
}

func IsValidDiseaseSearchInput(s string) bool {
	return diseaseRegex.MatchString(s)
}

func GetIdx(sequencechar byte) int {
	switch sequencechar {
	case 'A':
		return 0
	case 'T':
		return 1
	case 'C':
		return 2
	case 'G':
		return 3
	}
	return -1
}
