package rotate

func rotate(s []int, order int) {
	if len(s) < order || order <= 0 {
		return
	}

	tail := make([]int, order)
	copy(tail, s[0:order])
	copy(s, s[order:])
	for i := 0; i < len(tail); i++ {
		s[len(s)-order+i] = tail[i]
	}
}
