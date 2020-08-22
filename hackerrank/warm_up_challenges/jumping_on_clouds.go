package warmupchallenges

// JumpingOnClouds ...
func JumpingOnClouds(c []int) int32 {
	var jumps int32 = 0
	var i int = 0
	for i < len(c)-1 {

		if (i+2 < len(c)) && (c[i+2] == 0) {
			jumps++
			i += 2
		} else {
			jumps++
			i++
		}

	}
	return jumps
}
