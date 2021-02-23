package check

func isDight(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// check field match [0-9a-zA-z-_]+
func IsLegal(fieldName string) bool {
	for i := 0; i < len(fieldName); i++ {
		ch := fieldName[i]
		if !(isDight(ch) || isLetter(ch) || ch == '-' || ch == '_') {
			return false
		}
	}
	return true
}
