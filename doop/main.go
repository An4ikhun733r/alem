package main

import (
	"os"

	"github.com/01-edu/z01"
)

func prtZ01(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func atoi(str string) (int32, bool) {
	const maxInt32 = 1<<31 - 1

	var num int32
	negative := false

	if str[0] == '-' {
		negative = true
		str = str[1:]
	}

	if len(str) == 0 {
		return 0, false
	}

	for _, char := range str {
		if char < '0' || char > '9' {
			return 0, false
		}
		digit := int32(char - '0')

		// Check if adding the next digit will exceed the maximum value
		if num > maxInt32/10 || (num == maxInt32/10 && digit > maxInt32%10) {
			return 0, false
		}

		num = num*10 + digit
	}

	if negative {
		num = -num
	}

	return num, true
}

func intToString(num int32) string {
	if num == 0 {
		return "0"
	}

	var digits []byte
	negative := false

	if num < 0 {
		negative = true
		num = -num
	}

	for num > 0 {
		digit := byte(num%10 + '0')
		digits = append([]byte{digit}, digits...)
		num /= 10
	}

	if negative {
		digits = append([]byte{'-'}, digits...)
	}

	return string(digits)
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	num1, err1 := atoi(args[0])
	if err1 == false {
		return
	}
	operator := args[1]
	num2, err2 := atoi(args[2])
	if err2 == false {
		return
	}

	var result int32
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			prtZ01("No division by 0\n")
			return
		}
	case "%":
		if num2 != 0 {
			result = num1 % num2
		} else {
			prtZ01("No modulo by 0\n")
			return
		}
	default:
		return
	}

	prtZ01(intToString(result) + "\n")
}
