package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите пример: ")
	input, _ := reader.ReadString('\n')
	input = strings.ReplaceAll(input, " ", "")

	num1, operator, num2, isRoman := parseInput(input)

	result := calculate(num1, operator, num2, isRoman)

	if isRoman {
		fmt.Printf("Результат: %s\n", arabicToRoman(result))
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}

func parseInput(input string) (int, rune, int, bool) {
	operatorIndex := strings.IndexAny(input, "+-*/")
	if operatorIndex == -1 {
		panic("Выдача паники, так как строка не является математической операцией.")
	}

	operatorCount := strings.Count(input, "+") + strings.Count(input, "-") +
		strings.Count(input, "*") + strings.Count(input, "/")

	if operatorCount > 1 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	num1Str := strings.TrimSpace(input[:operatorIndex])
	num2Str := strings.TrimSpace(input[operatorIndex+1:])
	operator := rune(input[operatorIndex])

	isNum1Roman := isRomanNumeral(num1Str)
	isNum2Roman := isRomanNumeral(num2Str)

	if isNum1Roman != isNum2Roman {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}

	var num1, num2 int

	if isNum1Roman && isNum2Roman {
		num1 = romanToArabic(num1Str)
		num2 = romanToArabic(num2Str)
		if num1 > 10 || num2 > 10 {
			panic("Ошибка: Одно из введенных чисел больше 10")
		} else {
			return num1, operator, num2, true
		}
	} else if !isNum1Roman && !isNum2Roman {
		var err error
		num1, err = strconv.Atoi(num1Str)
		if err != nil {
			panic("Ошибка при конвертации первого числа: " + err.Error())
		}
		num2, err = strconv.Atoi(num2Str)
		if err != nil {
			panic("Ошибка при конвертации второго числа: " + err.Error())
		}
		if num1 > 10 || num2 > 10 {
			panic("Ошибка: Одно из введенных чисел больше 10")
		} else {
			return num1, operator, num2, false
		}
	}

	panic("Неизвестная ошибка при обработке ввода")
}

func isRomanNumeral(s string) bool {
	for _, r := range s {
		if !strings.ContainsRune("IVXLCDM", r) {
			return false
		}
	}
	return true
}

func calculate(num1 int, operator rune, num2 int, isRoman bool) int {
	switch operator {
	case '+':
		return num1 + num2
	case '-':
		if isRoman && num2 < num1 {
			return num1 - num2
		} else if isRoman {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		} else {
			return num1 - num2
		}
	case '*':
		return num1 * num2
	case '/':
		if num2 != 0 {
			return num1 / num2
		} else {
			panic("Ошибка: деление на ноль!")
		}
	default:
		panic("Неверный оператор!")
	}
}

func romanToArabic(roman string) int {
	romanNumerals := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	total := 0

	for i := 0; i < len(roman); i++ {
		if i+1 < len(roman) && romanNumerals[roman[i]] < romanNumerals[roman[i+1]] {
			total -= romanNumerals[roman[i]]
		} else {
			total += romanNumerals[roman[i]]
		}
	}

	return total
}

func arabicToRoman(num int) string {
	var romanPairs = []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	var roman strings.Builder
	for _, pair := range romanPairs {
		for num >= pair.Value {
			roman.WriteString(pair.Symbol)
			num -= pair.Value
		}
	}
	return roman.String()
}
