package utils

func AreEqual(a [16]byte, b [16]byte) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Contains(strings []string, s string) bool {
	for _, aString := range strings {
		if aString == s {
			return true
		}
	}
	return false
}

func Remove(strings []string, s string) []string {
	for i, aString := range strings {
		if aString == s {
			strings[i] = strings[0]
			strings = strings[1:]
			break
		}
	}
	return strings
}
