package warmupchallenges

// SockMerchant ...
func SockMerchant(arr []int) int {
	tempSocks := make(map[int]int)
	count := 0

	for _, v := range arr {
		if _, found := tempSocks[v]; found {
			count++
			delete(tempSocks, v)
		} else {
			tempSocks[v] = 1
		}
	}
	return count
}
