package logicalops

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

func inverse(b bool) bool{
	return !b
}


func and(x, y bool) bool{
	return x && y
}

func or(x, y bool) bool{
	return x || y
}

func deMorgan(x, y bool) (bool){
	return or(inverse(x), inverse(y))
}

