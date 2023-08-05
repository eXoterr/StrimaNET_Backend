package utils

func NilOrPanic(e error) {
	if e != nil {
		panic(e)
	}
}
