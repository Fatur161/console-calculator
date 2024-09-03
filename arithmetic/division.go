package arithmetic

func Divison(a int, b int) int {
	if b == 0 {
		panic("It cannot be divided by zero")
	}
	return a / b
}
