package utils

import (
	"regexp"
)

var dnaRegex *regexp.Regexp = regexp.MustCompile("^[ATCG]+$")
var diseaseRegex *regexp.Regexp = regexp.MustCompile(`^[\w-\s]*$`)
var inputSearchRegex *regexp.Regexp = regexp.MustCompile(`^(?:(?:(?:31(\s)+(?:Januari|Maret|Mei|Juli|Agustus|Oktober|Desember))|(?:(?:29|30)(\s)+(?:Januari|Maret|April|Mei|Juni|Juli|Agustus|September|Oktober|November|Desember)))(\s)*(\d{4})|^(?:29(\s)+(Februari)(\s)*(?:(\d{2}(?:[2468][02468]|[13579][26])|(?:[02468][048]|[13579][26])00)))|^(?:0?[1-9]|1\d|2[0-8])(\s)+(?:(?:Januari|Februari|Maret|April|Mei|Juni|Juli|Agustus|September)|(?:Oktober|November|Desember))(\s)+(\d{4})|(\s)*[\w-\s]+)(?:\s[\w-\s]+|\s*)$`)
var dateRegex *regexp.Regexp = regexp.MustCompile(`^(?:(?:(?:31(\s)+(?:Januari|Maret|Mei|Juli|Agustus|Oktober|Desember))|(?:(?:29|30)(\s)+(?:Januari|Maret|April|Mei|Juni|Juli|Agustus|September|Oktober|November|Desember)))(\s)*(\d{4})$|^(?:29(\s)+(Februari)(\s)*(?:(\d{2}(?:[2468][02468]|[13579][26])|(?:[02468][048]|[13579][26])00)))$|^(?:0?[1-9]|1\d|2[0-8])(\s)+(?:(?:Januari|Februari|Maret|April|Mei|Juni|Juli|Agustus|September)|(?:Oktober|November|Desember))(\s)+(\d{4}))$`)

func IsValidDNA(s string) bool {
	return dnaRegex.MatchString(s)
}

func IsValidDiseaseSearchInput(s string) bool {
	return diseaseRegex.MatchString(s)
}

func IsValidInputSearch(s string) bool {
	return inputSearchRegex.MatchString(s)
}

func IsValidDate(s string) bool {
	return dateRegex.MatchString(s)
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
