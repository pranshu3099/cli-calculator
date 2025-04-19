package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/chzyer/readline"
)

func getResult(num2 float64, opt string, total float64) float64 {

	switch opt {
	case "+":
		total = total + num2
	case "-":
		total = total - num2
	case "*":
		total = total * num2
	case "/":
		total = total / num2
	default:
		return total
	}

	return total
}

func getCalculation(input []string, total float64, opt string) float64 {
	for _, val := range input {
		if len(val) > 0 && unicode.IsDigit(rune(val[0])) {
			num, err := strconv.ParseFloat(val, 64)

			if err == nil {
				if total == 0 {
					if opt == "-" {
						total -= num
					} else {

						total += num
					}
				} else {
					total = getResult(num, opt, total)
				}

			}
		} else {
			opt = val
		}
	}

	return total

}

var rl *readline.Instance

func getInput(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := rl.Readline()
	return strings.ReplaceAll(input, " ", ""), err

}

func calculate() {

	for {

		input, err := getInput("")
		if err != nil {
			break
		}
		if input == "exit" {
			break
		}
		re := regexp.MustCompile(`\d+|[+\-*/]`)
		tokens := re.FindAllString(input, -1)
		result := getCalculation(tokens, 0, "")
		fmt.Println(result)
	}

}
