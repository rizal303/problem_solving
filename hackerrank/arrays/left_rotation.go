package arrays

// RotLeft ...
func RotLeft(a []int32, d int32) []int32 {
	i := int(d) % len(a)
	return append(a[i:], a[:i]...)
}
