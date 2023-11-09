package p0101

func isUnique(str string) bool {
	var count [26]bool
	for _, c := range str {
		if count[c-'a'] {
			return false
		}
		count[c-'a'] = true
	}
	return true
}

// isUniqueBit 使用位运算
func isUniqueBit(str string) bool {
	count := 0
	for _, c := range str {
		n := c - 'a'
		if (count>>n)&1 == 1 {
			return false
		}
		count |= 1 << n
	}
	return true
}
