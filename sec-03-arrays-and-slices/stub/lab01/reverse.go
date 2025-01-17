package main

// TODO 1 - implement 'string' reverse() function
func reverse(s string) string {

	sLen := len([]rune(s))
	currentIdx := sLen - 1
	revStrAsRune := make([]rune, sLen)

	for _, v := range s {
		revStrAsRune[currentIdx] = v
		currentIdx--
	}

	return string(revStrAsRune)

}
