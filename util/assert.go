package util

func Assert(a interface{}, b interface{}) bool {
	switch v := a.(type) {
	case int:
		switch j := b.(type) {
		case int:
			return assertInt(v, j)
		default:
			panic("type is not match! they need `int`")
		}

	case bool:
		switch j := b.(type) {
		case bool:
			return assertBool(v, j)
		default:
			panic("type is not match! they need `bool`")
		}
	case int64:
		switch j := b.(type) {
		case int64:
			return assertInt64(v, j)
		default:
			panic("type is not match! they need `int64`")
		}
	default:
		panic("not find assert !")
	}
}

func assertInt(a int, b int) bool {
	if a != b {
		return false
	}
	return true
}

func assertInt64(a int64, b int64) bool {
	if a != b {
		return false
	}
	return true
}

func assertBool(a bool, b bool) bool {
	if a == b {
		return true
	}
	return false
}
