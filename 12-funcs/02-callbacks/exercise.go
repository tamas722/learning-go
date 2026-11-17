package calculator

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

type OperationType int

const (
	Add OperationType = iota
	Subtract
	Multiply
)

func (op OperationType) String() string {
	switch op {
	case Add:
		return "Add"
	case Subtract:
		return "Subtract"
	case Multiply:
		return "Multiply"
	default:
		return "Unknown"
	}
}

type operationFunc func(float64, float64) float64

var operations = map[OperationType]operationFunc{
	Add:      func(a, b float64) float64 { return a + b },
	Subtract: func(a, b float64) float64 { return a - b },
	Multiply: func(a, b float64) float64 { return a * b },
}

func Calculate(op OperationType, a, b float64) float64 {
	if fn, ok := operations[op]; ok {
		return fn(a, b)
	}
	return 0
}
