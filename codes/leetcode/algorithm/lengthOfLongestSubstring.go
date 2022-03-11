package algorithm

func lengthOfLongestSubstring(s string) int {
	longest := 0
	repeat := make(map[string]struct{})
	for i := range s {
		currentLongest := 0
		for j := i; j < len(s); j++ {
			if _, exist := repeat[string(s[j])]; !exist {
				repeat[string(s[j])] = struct{}{}
				currentLongest += 1
			} else {
				repeat = map[string]struct{}{}
				break
			}
		}
		if currentLongest > longest {
			longest = currentLongest
		}

	}

	return longest
}
