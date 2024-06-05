package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var userInput string
var validOperators = "+-*/"
var validInputRomaNumeric = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
var validRomaNumeric = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
	"XXI": 21, "XXII": 22, "XXIII": 23, "XXIV": 24, "XXV": 25, "XXVI": 26, "XXVII": 27, "XXVIII": 28, "XXIX": 29, "XXX": 30,
	"XXXI": 31, "XXXII": 32, "XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36, "XXXVII": 37, "XXXVIII": 38, "XXXIX": 39, "XL": 40,
	"XLI": 41, "XLII": 42, "XLIII": 43, "XLIV": 44, "XLV": 45, "XLVI": 46, "XLVII": 47, "XLVIII": 48, "XLIX": 49, "L": 50,
	"LI": 51, "LII": 52, "LIII": 53, "LIV": 54, "LV": 55, "LVI": 56, "LVII": 57, "LVIII": 58, "LIX": 59, "LX": 60,
	"LXI": 61, "LXII": 62, "LXIII": 63, "LXIV": 64, "LXV": 65, "LXVI": 66, "LXVII": 67, "LXVIII": 68, "LXIX": 69, "LXX": 70,
	"LXXI": 71, "LXXII": 72, "LXXIII": 73, "LXXIV": 74, "LXXV": 75, "LXXVI": 76, "LXXVII": 77, "LXXVIII": 78, "LXXIX": 79, "LXXX": 80,
	"LXXXI": 81, "LXXXII": 82, "LXXXIII": 83, "LXXXIV": 84, "LXXXV": 85, "LXXXVI": 86, "LXXXVII": 87, "LXXXVIII": 88, "LXXXIX": 89, "XC": 90,
	"XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95, "XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100}

func main() {

	for {
		userInput, _ = reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		userInput = strings.ReplaceAll(userInput, " ", "")
		userInput = strings.ToUpper(userInput)

		operator := IsMath(userInput)
		operands := strings.Split(userInput, operator)
		oper1, oper2 := operands[0], operands[1]
		if IsValidRoma(oper1, oper2) {
			CalculRoma(oper1, oper2, operator)
			continue
		}
		if ok, a, b := IsValidArab(oper1, oper2); ok {
			CalculArab(a, b, operator)
			continue
		}

	}

}

func IsMath(ui string) string {
	var count, operator = 0, ""
	for _, op := range validOperators {
		if strings.Contains(ui, string(op)) {
			count = count + strings.Count(ui, string(op))
			operator = string(op)
		}
	}
	if count != 1 {
		panic("Невозможно вычислить")
	}
	return operator
}

func IsValidArab(str1 string, str2 string) (bool, int, int) {
	if str1 == "" || str2 == "" {
		panic("Невалидные операнды")
	}
	a, err := strconv.Atoi(str1)
	if err != nil {
		panic("Невалидные операнды")
	}
	b, err := strconv.Atoi(str2)
	if err != nil {
		panic("Невалидные операнды")
	}
	if !(a > 0 && a < 11) || !(b > 0 && b < 11) {
		panic("Невалидные операнды")
	}
	return true, a, b
}

func CalculArab(i1 int, i2 int, op string) {
	var ans int
	switch op {
	case "+":
		ans = i1 + i2
	case "-":
		ans = i1 - i2
	case "*":
		ans = i1 * i2
	case "/":
		ans = i1 / i2
	default:
		panic("Что-то пошло не так")
	}
	fmt.Println(ans)
}

func IsValidRoma(str1 string, str2 string) bool {
	return (validInputRomaNumeric[str1] != 0) && (validInputRomaNumeric[str2] != 0)
}

func CalculRoma(str1 string, str2 string, op string) {
	operand1, operand2, ans := validInputRomaNumeric[str1], validInputRomaNumeric[str2], 0
	switch op {
	case "+":
		ans = operand1 + operand2
	case "-":
		ans = operand1 - operand2
	case "*":
		ans = operand1 * operand2
	case "/":
		ans = operand1 / operand2
	default:
		panic("Что-то пошло не так")
	}
	if ans < 0 {
		panic("Результатом работы калькулятора с римскими числами могут быть только положительные числа")
	}
	key, ok := mapkey(validRomaNumeric, ans)
	if !ok {
		panic("Нет римского числа для полученного значения")
	}
	fmt.Println(key)
}

func mapkey(m map[string]int, value int) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}
