package calculo

import "errors"

var Operations = map[string]int{
	"suma":           1,
	"resta":          1,
	"division":       1,
	"multiplicacion": 1,
}

func Sumar(nums []int) int {
	suma := 0
	for _, num := range nums {
		suma += num
	}
	return suma
}

func Restar(nums []int) (int, error) {
	resta := 0
	for _, num := range nums {
		resta -= num
	}
	return resta, nil
}

func Dividir(num1, num2 int) (int, error) {
	if num2 != 0 {
		return num1 / num2, nil
	}
	return 0, errors.New("no se puede dividir entre 0")
}

func Multiplicar(nums []int) (int, error) {
	multiplicacion := nums[0]
	for _, num := range nums[1:] {
		multiplicacion *= num
	}
	return multiplicacion, nil
}
