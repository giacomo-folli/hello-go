package iteration

// Repeat takes in a character to repeat and a number of iterations
// Then it returns the repetition of character
func Repeat(character string, repeatedCount int) (word string) {
	if repeatedCount == 0 {
		repeatedCount = 5
	}
	for i := 0; i < repeatedCount; i++ {
		word += character
	}

	return
}
