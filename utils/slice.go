package utils

func DeleteFromRuneSlice(sli []rune, rm ...rune) []rune {
	var res []rune
	for _, r := range sli {
		if !ContainsRune(rm, r) {
			res = append(res, r)
		}
	}
	return res
}

func ContainsRune(sli []rune, v rune) bool {
	for _, r := range sli {
		if r == v {
			return true
		}
	}
	return false
}
