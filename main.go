package main

import (
	"fmt"
	"os"
	"strconv"

	"console_calculator/arithmetic"
	"console_calculator/utils"
)

func main(){
	var (
		command string
		firstNum int
		secondNum int
		result int
		exit bool
	)
	exit = true
	fmt.Println("Hello to console caclulator")
	for exit{
		fmt.Print("Please choose command:\n 1)+\n 2)-\n 3)*\n 4)/\n 0)exit\n")
		fmt.Fscan(os.Stdin, &command)
		i, err := strconv.Atoi(command)
		if err != nil{
			panic("unrecognized command")
		}
		if i > 0 && i < 5{
			firstNum, secondNum = utils.WriteNumber()
		}
		switch command{
		case "1":
			command = "+"
			result = arithmetic.Add(firstNum, secondNum)
		case "2":
			command = "-"
			result = arithmetic.Subtraction(firstNum, secondNum)
		case "3":
			command = "*"
			result = arithmetic.Multiplication(firstNum, secondNum)
		case "4":
			command = "/"
			result = arithmetic.Divison(firstNum, secondNum)
		case "0":
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			panic("unrecognized command")
		}
		fmt.Printf("%d %s %d = %d \n", firstNum, command, secondNum, result)
	}

}
