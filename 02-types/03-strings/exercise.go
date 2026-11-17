package strings

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

func multilineString() string {
	return "some\nmultiline\nstring"

}

func stringLen(s string) int {
	return len(s)
}

func trimFirstChar(s string) string {

	if s == "" {
		return s
	}
	return s[1:]
}

func trimLastChar(s string) string {
	if s == "" {
		return s
	}
	return s[:len(s)-1]
}

func swapFirstChar(s string) string{
	if s ==	""{	
		return s
	}

	temp := "A"
	temp2 := s[1:]

	return temp+temp2
}

func swapLastChar(s string) string{
	if s ==	""{	
		return s
	}

	temp := "A"
	temp2 := s[:len(s)-1]

	return temp2+temp
}

func prependChar(s string) string{
	return "A"+s
}

func appendChar(s string) string{
	return s+"A"
}

