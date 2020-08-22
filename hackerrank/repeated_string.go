package hackerrank

import (
	"strings"
)

// RepeatedString ...
func RepeatedString(s string, n int64) int64 {
	k := strings.Count(s, "a") * (int(n) / len(s))
	k += strings.Count(s[:int(n)%len(s)], "a")
	return int64(k)
}
