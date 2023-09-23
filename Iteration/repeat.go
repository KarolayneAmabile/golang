package iteration

const repeatCount = 5

func Repeat(character string) string { // declare a function named repeat that takes an input argument character of type string and returns a string as the result
	var repeated string // declare a variable/name/data type
	for i := 0; i < repeatCount; i++ {
		repeated += character // the add and assignment operator
	}
	return repeated
}

// we've been using := to declare and initializing variables, however, := is
// simply short hand for both steps
