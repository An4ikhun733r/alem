package piscine

func isAlphanumeric(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

func Capitalize(s string) string {
	bytes := []byte(s)
	capitalized := make([]byte, len(bytes))

	shouldCapitalize := true
	for i, ch := range bytes {
		if isAlphanumeric(ch) {
			if shouldCapitalize {
				if ch >= 'a' && ch <= 'z' {
					ch -= 32
				}
				capitalized[i] = ch
			} else {
				if ch >= 'A' && ch <= 'Z' {
					ch += 32
				}
				capitalized[i] = ch
			}
			shouldCapitalize = false
		} else {
			capitalized[i] = ch
			shouldCapitalize = true
		}
	}

	return string(capitalized)
}
