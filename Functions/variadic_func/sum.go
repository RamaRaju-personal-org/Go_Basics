package variadic_func

func Sumup(num ...int) int { // ...int passing n number of values of type int
	sum := 0

	for _, val := range num {
		sum += val
	}
	return sum
}
