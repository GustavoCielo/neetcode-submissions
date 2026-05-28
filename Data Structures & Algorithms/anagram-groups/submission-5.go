// import "slices"

func groupAnagrams(strs []string) [][]string {
	output := [][]string{}
	if len(strs) == 1 {
		output = append(output, strs)
		return output
	}
	seen := make(map[int]bool)
	for i := 0; i < len(strs); i++ {
		freshSlice := []string{}
		if !seen[i] {
			seen[i] = true
			freshSlice = append(freshSlice, strs[i])
		}
		for j := i + 1; j < len(strs); j++ {
			if checkAnagrams(strs[i], strs[j]) {
				if !seen[j] { 
					freshSlice = append(freshSlice, strs[j])
					seen[j] = true
				}
			}
		}
		if len(freshSlice) == 0 {
			continue
		}
		fmt.Println(freshSlice)
		output = append(output, freshSlice)
	}
	return output
}

func checkAnagrams(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := make(map[rune]int)
	for _, v := range s {
		counts[v]++
	}
	for _, v := range t {
		counts[v]--
		if counts[v] < 0 {
			return false
		}
	}
	return true
}