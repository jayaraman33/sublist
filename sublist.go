package sublist

type Relation string

func Sublist(a, b []int) Relation {
	
	switch {
	case equal(a, b):
		return "equal"
	case sublist(a, b):
		return "sublist"
	case sublist(b, a):
		return "superlist"
	default:
		return "unequal"
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func sublist(a, b []int) bool {
	for i := range b {
		if i+len(a) > len(b) {
			return false
		}
		if equal(a, b[i:len(a)+i]) {
			return true
		}
	}
	return false
}
