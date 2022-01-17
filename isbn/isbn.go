package isbn

// ValidISBN10 validates the string in isbn format.
func ValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}

	var s, n int

	for i, c := range isbn {
		if i == len(isbn)-1 && (c == 'X' || c == 'x') {
			n = 10
		} else if c < '0' || c > '9' {
			return false
		} else {
			n = int(c - '0')
		}
		s += n * (i + 1)
	}
	return s%11 == 0
}
