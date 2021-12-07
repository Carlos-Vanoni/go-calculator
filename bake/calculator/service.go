package main

import "fmt"

var history []string

func save(result string) {
	history = append(history, result)
}

func getHistory() []string {
	return history
}

type calculator struct {
	num1 int
	num2 int
}

func (calculator calculator) sum() string {
	var result = calculator.num1 + calculator.num2
	save(fmt.Sprint(calculator.num1) + "+" + fmt.Sprint(calculator.num2) + "=" + fmt.Sprint(result))
	return fmt.Sprint(result)
}

func (calculator calculator) sub() string {
	var result = calculator.num1 - calculator.num2
	save(fmt.Sprint(calculator.num1) + "-" + fmt.Sprint(calculator.num2) + "=" + fmt.Sprint(result))
	return fmt.Sprint(result)
}

func (calculator calculator) mul() string {
	var result = calculator.num1 * calculator.num2
	save(fmt.Sprint(calculator.num1) + "*" + fmt.Sprint(calculator.num2) + "=" + fmt.Sprint(result))
	return fmt.Sprint(result)
}

func (calculator calculator) div() string {
	var result = calculator.num1 / calculator.num2
	save(fmt.Sprint(calculator.num1) + "/" + fmt.Sprint(calculator.num2) + "=" + fmt.Sprint(result))
	return fmt.Sprint(result)
}
