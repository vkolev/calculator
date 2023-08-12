package main

import (
	"calculator"
	"fmt"
	"os"
	"strconv"
)

func printHelp() {
	fmt.Println("Usage: calculator [command] <value1> <value2:optional>")
	fmt.Println()
	fmt.Println("Available commands")
	fmt.Println("add\t\tAdd the two values")
	fmt.Println("sub\t\tSubstract second value from first value")
	fmt.Println("mul\t\tMultiply the two values")
	fmt.Println("div\t\tDivide first value by second value")
	fmt.Println("sqrt\t\tGet Square root of a given value")
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 || argsWithoutProg[0] == "help" {
		printHelp()
		return
	}
	var result float64
	a, err := strconv.ParseFloat(argsWithoutProg[1], 64)
	if err != nil {
		fmt.Println("Invalid value for 'a'")
		return
	}
	var b float64
	if len(argsWithoutProg) == 3 {
		b, _ = strconv.ParseFloat(argsWithoutProg[2], 64)
	}

	switch command := argsWithoutProg[0]; command {
	case "add":
		result = calculator.Add(a, b)
		fmt.Println(result)
	case "sub":
		result = calculator.Substract(a, b)
		fmt.Println(result)
	case "mul":
		result = calculator.Multiply(a, b)
		fmt.Println(result)
	case "div":
		result, err = calculator.Divide(a, b)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	case "sqrt":
		result, err = calculator.Sqrt(a)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	default:
		printHelp()
	}
}
