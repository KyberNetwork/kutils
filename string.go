package kutils

import (
	"strings"
)

// SplitListElem split each element in a string list with the given separator
func SplitListElem(lst []string, sep string) []string {
	ret := make([]string, 0, len(lst))
	for _, elem := range lst {
		ret = append(ret, strings.Split(elem, sep)...)
	}
	return ret
}
