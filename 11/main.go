//Реализовать пересечение двух неупорядоченных множеств.
package main

import "fmt"

func intersection1(x []int, y []int) []int {
	// set - коллекция, которая реализует основные математические операции над множествами
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	// отмечаем элементы тру (игнорируем повторы)
	for _, val := range x {
		set1[val] = true
	}

	for _, val := range y {
		set2[val] = true
	}

	// берем большее множесто и итерируемся по нему
	// а в меньшем проверяем на вхождение
	// если да, добавляем в слайс

	out := make([]int, 0, len(x))
	for i := range set2 {
		if set1[i] {
			out = append(out, i)
			fmt.Println(cap(out))
		}
	}
	return out
}
func main() {
	firstArray := []int{5, 7, 1, 2, 3}
	secondArray := []int{9, 8, 1, 4, 2, 5, 6, 3, 4, 5, 23, 4, 3, 2}
	c := intersection1(firstArray, secondArray)
	fmt.Println(c)
}
