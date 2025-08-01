package main

// ChatGPT
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

// Сам написал, да)))
func backspaceCompare(s string, t string) bool {
	pntS := len(s) - 1
	pntT := len(t) - 1
	shiftS := 0
	shiftT := 0
	for pntS >= 0 || pntT >= 0 {
		for pntS >= 0 && (shiftS > 0 || s[pntS] == '#') {
			if s[pntS] == '#' {
				shiftS++
			} else {
				shiftS--
			}
			pntS--
		}
		for pntT >= 0 && (shiftT > 0 || t[pntT] == '#') {
			if t[pntT] == '#' {
				shiftT++
			} else {
				shiftT--
			}
			pntT--
		}
		if (pntS < 0 || pntT < 0) && (pntS >= 0 || pntT >= 0) {
			return false
		}
		if pntS >= 0 && pntT >= 0 && s[pntS] != t[pntT] {
			return false
		}
		pntS--
		pntT--
	}
	return true
}

func main() {

}
