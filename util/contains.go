package util

func StringSliceContains(x string, other []string) bool {
	for _, o := range other {
		if o == x {
			return true
		}
	}
	return false
}
