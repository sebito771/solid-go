package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {

		sum += number
	}
	return sum

}

func Sumall(NumberstoSum ...[]int) []int {
	lengthofnumbers := len(NumberstoSum)
	Sums := make([]int, lengthofnumbers)
	for i, number := range NumberstoSum {
		Sums[i] = Sum(number)

	}
	return Sums
}
func Sumalltails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 { // Verificar si está vacío
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] // Solo obtener el tail si no está vacío
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
