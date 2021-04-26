package syntax

func FactorialFor(n int) int {
	fac := 1
	for i := 1; i <= n; i++ {
		fac *= i
	}
	return fac
}

func FactorialWhile(n int) int {
	fac := n
	for n > 1 {
		n--
		fac *= n
	}
	return fac
}

func Hello(name string) string {
	if name == "miku" {
		return "Hello " + name
	} else if name == "luka" {
		return "Bonjour " + name
	} else {
		return "Chow " + name
	}
}

func IfWithStatement() string {
	if v := 1 + 2; v < 1 {
		return "nope"
	}
	return "yes"
}

func RandomizeString(value string, seed int) string {
	switch l := len(value) * seed % 4; l {
	// case clause can be an expression
	case isOdd(l):
		return randomizeOddString(value)
	case 0:
		return value
	default:
		return randomizeEvenString(value)
	}
}

func isOdd(value int) int {
	if value%2 == 0 {
		return value
	}
	return -value
}

func randomizeOddString(value string) string {
	return value
}

func randomizeEvenString(value string) string {
	return value
}

func SwitchWithoutCondition(value int) int {
	switch {
	case value > 0:
		return 1
	case value < 0:
		return -1
	default:
		return 0
	}
}

func DeferExecuteAfterFunctionReturns() {
	defer isOdd(1)
}
