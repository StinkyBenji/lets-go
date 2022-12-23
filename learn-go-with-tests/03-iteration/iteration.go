package iteration

// in Go, there is only for, no while, do, until
const repeatCount = 6

func Iteration(character string) string {
	//  := to declare and initialize var
	// e.g. f := "apple" ; var f string = "apple"
	// var declares 1 or more variables, zero-valued

	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
