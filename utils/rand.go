package utils

import "math/rand"

func RumpleRuneArray(rarr []rune) []rune {
	for i := len(rarr) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		rarr[i], rarr[num] = rarr[num], rarr[i]
	}

	return rarr
}

func RandomSelectRuneArray(rarr []rune, n int) []rune {
	return RumpleRuneArray(rarr)[:n]
}

func RandomAlphabetMappingTable() map[rune]rune {
	var mappingKeys []rune
	for key := 'A'; key <= 'Z'; key++ {
		mappingKeys = append(mappingKeys, key)
	}
	val := RumpleRuneArray(mappingKeys)
	mappingTable := make(map[rune]rune)
	for idx, key := 0, 'A'; key <= 'Z'; idx, key = idx+1, key+1 {
		mappingTable[key] = val[idx]
	}
	return mappingTable
}
