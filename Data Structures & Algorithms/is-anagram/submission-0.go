func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	s1 := map[string]int{}
	for _, v := range strings.Split(s, "") {
		if _, ok := s1[v]; ok {
			s1[v] += 1
		} else {
			s1[v] = 1
		}
	}
	s2 := map[string]int{}
	for _, v := range strings.Split(t, "") {
		if _, ok := s2[v]; ok {
			s2[v] += 1
		} else {
			s2[v] = 1
		}
	}
	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
