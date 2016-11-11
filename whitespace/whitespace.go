package whitespace

//IsWhitespace checks if a rune is whitespace
func IsWhitespace(r rune) bool {
	switch r {
	case '\u0009':
		return true
	case '\u000A':
		return true
	case '\u000B':
		return true
	case '\u000C':
		return true
	case '\u000D':
		return true
	case '\u0020':
		return true
	case '\u0085':
		return true
	case '\u00A0':
		return true
	case '\u1680':
		return true
	case '\u2000':
		return true
	case '\u2001':
		return true
	case '\u2002':
		return true
	case '\u2003':
		return true
	case '\u2004':
		return true
	case '\u2005':
		return true
	case '\u2006':
		return true
	case '\u2007':
		return true
	case '\u2008':
		return true
	case '\u2009':
		return true
	case '\u200A':
		return true
	case '\u2028':
		return true
	case '\u2029':
		return true
	case '\u202F':
		return true
	case '\u205F':
		return true
	case '\u3000':
		return true
	}

	return false
}

//ContainsWhitespace checks if str contain whitespaces
func ContainsWhitespace(str string) bool {
	for _, c := range str {
		if IsWhitespace(c) {
			return true
		}
	}
	return false
}

//OnlyWhitespace checks if a string is only whitespace
func OnlyWhitespace(str string) bool {
	for _, c := range str {
		if !IsWhitespace(c) {
			return false
		}
	}
	return true
}
