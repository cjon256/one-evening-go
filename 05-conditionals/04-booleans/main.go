package conditionals

func In20thCentury(year int) bool {
	if 1900 < year && year < 2001 {
		return true
	}
	return false
}
