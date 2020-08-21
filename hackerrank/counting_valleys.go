package hackerrank

// CountingValleys ...
func CountingValleys(s string) int32 {
	level := 0
	var valleys int32 = 0
	for _, v := range s {
		if string(v) == "U" {
			level++
			if level == 0 {
				valleys++
			}
		} else {
			level--
		}
	}
	return valleys
}
