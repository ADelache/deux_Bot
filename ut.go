package ut

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func countInArray(array []string, what string) int {
	count := 0
	for i := 0; i < len(array); i++ {
		if array[i] == what {
			count++
		}
	}
	return count
}

func countInArrayI(array []int, what int) int {
	count := 0
	for i := 0; i < len(array); i++ {
		if array[i] == what {
			count++
		}
	}
	return count
}
