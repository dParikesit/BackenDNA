package utils

import "fmt"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CalculateLevenshteinDist(dna string, disease string) float32 {
	var ratio float32 = 0
	var count int = 0
	var lenDisease = len(disease)
	var lenDna = len(dna)
	var totallength int = lenDisease + lenDna

	var matrix = make([][]int, lenDisease+1)
	for i := range matrix {
		matrix[i] = make([]int, lenDna+1)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// Fill matrix with zero
			matrix[i][j] = 0
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			matrix[i][0] = i
			matrix[0][j] = j
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if disease[i-1] == dna[j-1] {
				count = 0
			} else {
				count = 2
			}
			matrix[i][j] = Min(matrix[i-1][j]+1, matrix[i][j-1]+1)
			matrix[i][j] = Min(matrix[i-1][j-1]+count, matrix[i][j])
		}
	}

	var lastElement int = matrix[lenDisease][lenDna]
	ratio = float32((totallength - lastElement)) / float32(totallength)
	fmt.Println(ratio)
	return ratio
}

func SimiliarityMatch(dna string, disease string) (bool, float32) {
	// use levenshtein distance algorithm
	lenDisease := len(disease)
	lenDna := len(dna)

	if lenDisease > lenDna {
		return false, 0.0
	}

	var temp float32
	var biggestRatio float32
	biggestRatio = -0.1
	for i := 0; i < lenDna-lenDisease; i++ {
		first10 := dna[i : i+lenDisease]
		temp = CalculateLevenshteinDist(first10, disease)

		if temp > biggestRatio {
			biggestRatio = temp
		}
	}

	return (biggestRatio >= 0.8), biggestRatio
}
