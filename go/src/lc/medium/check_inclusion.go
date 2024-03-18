package medium

import (
	"slices"
)

// checkInclusion (567 | Medium)
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Map, s2Map := [26]int{}, [26]int{}
	for i := range s1 {
		s1Map[s1[i]-'a']++
		s2Map[s2[i]-'a']++
	}
	for j := 0; j < len(s2)-len(s1); j++ {
		if slices.Equal(s1Map[:], s2Map[:]) {
			return true
		}
		s2Map[s2[len(s1)+j]-'a']++
		s2Map[s2[j]-'a']--
	}
	return slices.Equal(s1Map[:], s2Map[:])
}
