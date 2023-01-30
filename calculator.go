package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	read := bufio.NewReader(os.Stdin)

	roman := map[string]string{
		"I": "1", "II": "2",
		"III": "3", "IV": "4",
		"V": "5", "VI": "6",
		"VII": "7", "VIII": "8",
		"IX": "9", "X": "10",
	}

	fmt.Println("Привет. Перед тобой Арабо-Римский калькулятор.")

letsgo:
	for {

		fmt.Println("Введи два операнда и один оператор через пробел (пример: a + b, a - b, a * b, a / b):")
		inp, _ := read.ReadString('\n')
		inp = strings.TrimSpace(inp)
		inp = strings.ToUpper(inp)

		oper := strings.Split(inp, " ")

		if len(oper) != 3 {
			fmt.Println("Ошибка: так как формат математической операции не удовлетворяет заданию — два операнда и один оператор через пробел (+, -, /, *)")
			continue
		}

		roman1 := romck(oper[0], roman)
		roman2 := romck(oper[2], roman)
		arab1 := arabck(oper[0], roman)
		arab2 := arabck(oper[2], roman)

		switch {
		case roman1 == true && roman2 == true:
			d1, _ := strconv.Atoi(roman[oper[0]])
			operation := oper[1]
			d2, _ := strconv.Atoi(roman[oper[2]])
			result := count(d1, d2, operation)
			if result == -1 {
				fmt.Println("Ошибка: так как формат математической операции не удовлетворяет заданию — два операнда и один оператор через пробел (+, -, /, *)")
			} else if result <= 0 {
				fmt.Println("Ошибка: так как в римской системе нет отрицательных чисел.")
			} else {
				fmt.Println("Результат:", resultToRoman(result))
			}

		case arab1 == true && arab2 == true:
			d1, _ := strconv.Atoi(oper[0])
			operation := oper[1]
			d2, _ := strconv.Atoi(oper[2])
			result := countforarab(d1, d2, operation)
			if result == -1 {
				fmt.Println("Ошибка: так как формат математической операции не удовлетворяет заданию — два операнда и один оператор через пробел (+, -, /, *)")
			} else {
				fmt.Println("Результат:", result)
			}

		case (arab1 || arab2) && (roman1 || roman2):
			fmt.Println("Ошибка: так как используются одновременно разные системы счисления.")
			break letsgo

		case arab1 || arab2:
			fmt.Println("Ошибка: так как строка не является математической операцией.")
			break letsgo

		case roman1 || roman2:
			fmt.Println("Ошибка: так как строка не является математической операцией.")
			break letsgo

		default:
			fmt.Println("Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более.")
			break letsgo
		}
	}
}

func romck(d string, roman map[string]string) bool {

	for i, _ := range roman {
		if d == i {
			return true
		}
	}
	return false
}

func arabck(d string, roman map[string]string) bool {

	for _, v := range roman {
		if d == v {
			return true
		}
	}
	return false
}

func count(d1, d2 int, operation string) int {

	switch operation {
	case "+":
		return d1 + d2
	case "-":
		return d1 - d2
	case "*":
		return d1 * d2
	case "/":
		return d1 / d2
	}
	return -1
}

func countforarab(d1, d2 int, operation string) float64 {

	switch operation {
	case "+":
		return float64(d1) + float64(d2)
	case "-":
		return float64(d1) - float64(d2)
	case "*":
		return float64(d1) * float64(d2)
	case "/":
		return float64(d1) / float64(d2)
	}
	return -1
}

func resultToRoman(num int) string {

	romnum := map[int]string{
		1: "I", 2: "II",
		3: "III", 4: "IV",
		5: "V", 6: "VI",
		7: "VII", 8: "VIII",
		9: "IX", 10: "X",
		50: "L", 100: "C",
	}
	switch {
	case num <= 10 || num == 100:
		return romnum[num]
	case num <= 39:
		d1, d2 := difNum(num)
		output1 := strings.Repeat(romnum[10], d1)
		return output1 + romnum[d2]
	case num <= 49:
		_, d2 := difNum(num)
		return romnum[10] + romnum[50] + romnum[d2]
	case num <= 89:
		d1, d2 := difNum(num)
		output1 := strings.Repeat(romnum[10], d1-5)
		return romnum[50] + output1 + romnum[d2]
	case num <= 99:
		_, d2 := difNum(num)
		return romnum[10] + romnum[100] + romnum[d2]
	}
	return "0"
}

func difNum(d int) (int, int) {

	tostr := strconv.Itoa(d)
	ls := strings.Split(tostr, "")
	d1, _ := strconv.Atoi(ls[0])
	d2, _ := strconv.Atoi(ls[1])
	return d1, d2
}
