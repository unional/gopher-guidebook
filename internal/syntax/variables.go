package syntax

// declare requires type
var C, Python, Java bool

// declare with initial values
var I, J int = 1, 2

// declares multiple types
var (
	X, Y, Z int
	P, Q, R string
)

// lowercase for local variables and functions
var localVariable = 1

// function with return value must declare return type
func VariableDeclaration() int {
	// declare with implicit type
	a := 1
	return localVariable + a
}

const Pi = 3.14

const (
	Apple  = "red"
	Orange = "orange"
	Banana = "yellow"
)
