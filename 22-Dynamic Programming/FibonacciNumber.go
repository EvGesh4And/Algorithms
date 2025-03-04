package main

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	prev2, prev1 := 0, 1 // Запоминаем только два последних числа
	for i := 2; i <= n; i++ {
		curr := prev1 + prev2
		prev2, prev1 = prev1, curr // Обновляем значения
	}

	return prev1
}
