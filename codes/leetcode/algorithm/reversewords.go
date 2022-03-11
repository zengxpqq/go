package algorithm

import "strings"

func reverseWords(s string) string {
	splitStr := strings.Split(s, " ")

	reverse := func(str string) string {
		sub := []byte(str)
		for i, j := 0, len(sub)-1; i < len(sub)/2; i++ {
			sub[i], sub[j-i] = sub[j-i], sub[i]
		}
		return string(sub)
	}

	for i, sub := range splitStr {
		splitStr[i] = reverse(sub)
	}
	return strings.Join(splitStr, " ")
}
