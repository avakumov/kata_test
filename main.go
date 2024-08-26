package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func removeSpaces(text string) string {
	res := strings.ReplaceAll(text, "\t", "")
	res = strings.ReplaceAll(res, "\n", "")
	return strings.ReplaceAll(res, " ", "")
}

func isRange(number int) bool {
	if number < 1 || number > 10 {
		return false
	}
	return true
}

func isRoman(number string) bool {
	for _, r := range number {
		if !(r == 'I' || r == 'V' || r == 'X') {
			return false
		}
	}

	return true
}

func romanToInt(s string) int {
	res := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		tmp := m[s[i]]
		sign := 1
		if tmp < last {
			sign = -1
		}
		res += sign * tmp
		last = tmp
	}
	return res
}

func intToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func handleOperation(number1 int, operation string, number2 int) int {
	//check if number is in range
	if isRange(number1) && isRange(number2) {
		switch operation {
		case "+":
			return number1 + number2
		case "-":
			return number1 - number2
		case "*":
			return number1 * number2
		case "/":
			return number1 / number2

		}
	} else {
		panic("number out of range")
	}
	return 0

}

func handleWithRoman(first string, operation string, second string) string {
	number1 := romanToInt(first)
	number2 := romanToInt(second)
	result := handleOperation(number1, operation, number2)

	if result < 1 {
		panic("there are no negative Roman numbers")
	}
	return intToRoman(result)
}

func handleCommand(str string) string {
	text := removeSpaces(str)
	switch text {
	case "exit":
		os.Exit(0)
	case "help":
		fmt.Println("type exit to exit")
		fmt.Println("examples: 1 + 6, V - II")
	default:
		{
			r, _ := regexp.Compile(`(\+|\-|\*|\/)`)
			operations := r.FindAllString(text, -1)
			if operations == nil {
				panic("invalid operation. Use +, -, *, or /")
			}
			if len(operations) != 1 {
				panic("invalid operation. Use only one operation")
			}

			values := strings.Split(text, operations[0])

			//operation with roman numbers
			if isRoman(values[0]) && isRoman(values[1]) {
				return handleWithRoman(values[0], operations[0], values[1])

			//operation with arabic numbers
			} else {
				number1, err := strconv.Atoi(values[0])
				if err != nil {
					panic("first number is not a integer")
				}

				number2, err := strconv.Atoi(values[1])
				if err != nil {
					panic("second number is not a integer")
				}

				result := handleOperation(number1, operations[0], number2)
				return strconv.Itoa(result)
			}

		}
	}
	return ""
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("type help to get help")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		var result = handleCommand(text)
		fmt.Println(result)

	}

}
