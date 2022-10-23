package utils

func RuneMapToStringMap(m map[rune]rune) map[string]string {
	res := make(map[string]string)
	for k, v := range m {
		res[string(k)] = string(v)
	}
	return res
}

func RuneSliceToStringSlice(s []rune) []string {
	var res []string
	for _, c := range s {
		res = append(res, string(c))
	}
	return res
}
