package utils

func BMMatch(dna string, disease string) bool {
	lenDisease := len(disease)
	lenDna := len(dna)
	L := GetLastOccurence(disease)

	if lenDisease > lenDna {
		return false
	}

	var i, j int
	i = lenDisease - 1
	j = lenDisease - 1

	for {
		if disease[j] == dna[i] {
			if j == 0 {
				// After matching all the characters, if the last character of the disease is matched, then return true
				return true
			} else {
				i--
				j--
			}

		} else {
			var idxInLast int = GetIdx(dna[i])
			var lastOcc int = L[idxInLast]
			i = i + lenDisease - Min(j, lastOcc+1)
			j = lenDisease - 1
		}

		if i > lenDna-1 {
			break
		}
	}
	return false
}

func GetLastOccurence(disease string) []int {
	// get last occurence for ATCG in disease
	lenDisease := len(disease)
	var L = make([]int, 4)
	for i := 0; i < 4; i++ {
		L[i] = -1
	}

	for i := 0; i < lenDisease; i++ {
		switch disease[i] {
		case 'A':
			L[0] = i
		case 'T':
			L[1] = i
		case 'C':
			L[2] = i
		case 'G':
			L[3] = i
		}
	}
	return L
}
