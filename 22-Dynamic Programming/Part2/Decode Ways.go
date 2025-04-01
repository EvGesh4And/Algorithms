package main

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	prev2 := 1 // Для пустой строки
	prev1 := 1 // Для первого символа

	for i := 1; i < len(s); i++ {
		curr := 0
		currentDigit := s[i] - '0'
		prevDigit := s[i-1] - '0'
		twoDigits := prevDigit*10 + currentDigit

		// Если текущая цифра не '0', можно декодировать её отдельно
		if currentDigit > 0 {
			curr += prev1
		}

		// Если предыдущая и текущая цифры образуют число от 10 до 26
		if prevDigit > 0 && twoDigits <= 26 {
			curr += prev2
		}

		// Если декодирование невозможно
		if curr == 0 {
			return 0
		}

		// Обновляем prev2 и prev1 для следующего шага
		prev2, prev1 = prev1, curr
	}

	return prev1
}
