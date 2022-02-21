package main

// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.

/*
	var justString string

	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}

	func main() {
		someFunc()
	}
*/

/*
	Первый раз вижу объявленную переменную вне функции:D
	Скорее всего - она глобальная. Скорее всего, как и везде, глобальные переменные это плохо.

	https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d
	Насколько я смог узнать отсюда, в Go есть два вида "областей памяти". Stack и Heap.
	Хип - общий для всех горутин. Это значит, шо доступ к нему будет у всех.

	Так как строка - это массив байтов, то строка в utf != строка на каком нибудь японском языке.
	Не можем быть уверены, что получим ровно то количество символов, которое хотим

	Если слайс имеет 100 значений, а потом мы из него удалим 50, размер слайса не изменится.
	Он останется 100, просто 50 из них будут нули. В нашем примере мы ссылаемся на большой слайс.
	Следовательно, наверное, храним в памяти и его. Но не точно.
*/

// хотя бы убираем глобальную переменную
/*
	func someFunc() []byte {
		v := createHugeString(1 << 10)
		justString := append([]byte{}, v[:100]...)
		return justString
	}

	func main() {
		str := someFunc()
	}
*/