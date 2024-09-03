package utils

import (
	"fmt"
	"os"
)

func WriteNumber() (num1 int, num2 int){
	fmt.Print("Enter the first number: ")
	fmt.Fscan(os.Stdin, &num1)
	fmt.Print("Enter the second number: ")
	fmt.Fscan(os.Stdin, &num2)
	return
}
