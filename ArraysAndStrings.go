package problems

// StringHasAllUniqueCharacters determines if the passed string has all unique characters
func StringHasAllUniqueCharacters(input string) bool {
	keymap := make(map[rune]bool)
	for _, c := range input {
		if _, ok := keymap[c]; ok {
			return false
		}
		keymap[c] = true
	}

	return true
}

// StringReverseNullTerminated reverses a null-terminated string.
func StringReverseNullTerminated(input string) string {
	// Let's create a "null-terminated" string.
	// Convert the string to a list of runes and then append a 0 to act as a null character.
	runesList := append([]rune(input), 0)

	// Because this is "null-terminated", let's assume that `len` isn't available to us and calculate it.
	length := 0
	currentRune := runesList[length]
	for currentRune != 0 {
		length++
		currentRune = runesList[length]
	}

	// Reverse in-place
	start := 0
	end := length - 1

	for start < end {
		temp := runesList[start]
		runesList[start] = runesList[end]
		runesList[end] = temp
		start++
		end--
	}

	// Strip out the null termination and convert back to string
	return string(runesList[:len(runesList)-1])
}
