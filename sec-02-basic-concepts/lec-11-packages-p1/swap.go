package main

// parameters both int
func swap(s, t int) (int, int) {
	return t, s
}

// can also specify which values we cant to return
func swap2(s, t int) (x int, y int) { // function signature
	// function body starts here
	x = t
	y = s
	// return 10, 25 --> something like this is still valid
	return
}
