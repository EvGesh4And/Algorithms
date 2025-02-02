package main

func backspaceCompare(s string, t string) bool {
	pS, pT := len(s)-1, len(t)-1

	for pS >= 0 || pT >= 0 {
		// Найти следующий значащий символ в s
		skip := 0
		for pS >= 0 {
			if s[pS] == '#' {
				skip++
			} else if skip > 0 {
				skip--
			} else {
				break
			}
			pS--
		}

		// Найти следующий значащий символ в t
		skip = 0
		for pT >= 0 {
			if t[pT] == '#' {
				skip++
			} else if skip > 0 {
				skip--
			} else {
				break
			}
			pT--
		}

		// Сравниваем символы
		if pS >= 0 && pT >= 0 && s[pS] != t[pT] {
			return false
		}

		// Если один указатель остался, а другой вышел за границы
		if (pS >= 0) != (pT >= 0) {
			return false
		}

		pS--
		pT--
	}

	return true
}


func main() {

}
