package main

import "fmt"

// Map принимает функцию action и срез данных data.
// Функция Map должна применить функцию action к каждому элементу среза data и вернуть новый срез с результатами
func Map(data []int, action func(int) int) []int {
	if data == nil {
		return nil
	}

	if len(data) == 0 {
		return []int{}
	}

	result := make([]int, len(data))

	for i := 0; i < len(data); i++ {
		result[i] = action(data[i])
	}

	return result
}

// Filter принимает функцию action и срез данных data.
// Функция Filter должна вернуть новый срез, содержащий только те элементы data, для которых функция action возвращает true
func Filter(data []int, action func(int) bool) []int {
	if data == nil {
		return nil
	}

	if len(data) == 0 {
		return []int{}
	}

	result := make([]int, 0, cap(data))

	for i := 0; i < len(data); i++ {
		if action(data[i]) {
			result = append(result, data[i])
		}
	}

	return result
}

// Reduce принимает функцию action (функцию двух аргументов), срез данных data и начальное значение initial.
// Функция Reduce должна применить функцию action к каждому элементу data и начальному значению initial, накапливая результат.
func Reduce(data []int, initial int, action func(int, int) int) int {
	result := initial

	for i := 0; i < len(data); i++ {
		result = action(result, data[i])
	}

	return result
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6}

	//resultMap := Map(sl, func(el int) int {
	//	fmt.Printf("el: %d\n", el)
	//	return el * 2
	//})
	//fmt.Println(resultMap)
	//
	//fmt.Println("---")

	//resultFilter := Filter(sl, func(el int) bool {
	//	return el%2 == 0
	//
	//})
	//fmt.Printf("resultFilter: %d\n", resultFilter)
	//
	//fmt.Println("---")

	reduceResult := Reduce(sl, 0, func(acc int, curr int) int {
		return acc + curr
	})

	fmt.Printf("reduceResult: %d\n", reduceResult)

	fmt.Println("---")
}
