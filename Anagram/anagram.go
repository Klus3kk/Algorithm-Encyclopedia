func Anagrams(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	anagramMap := make(map[string]int)

	for i := 0; i < len(s1); i++ {
		anagramMap[string(s1[i])]++
	}

	for i := 0; i < len(s2); i++ {
		anagramMap[string(s2[i])]--
	}

	for i := 0; i < len(s1); i++ {
		if anagramMap[string(s1[i])] != 0 {
			return false
		}
	}
	return true
}