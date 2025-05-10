func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	res := make([][]string, 0, len(searchWord))

	prefix := ""
	start := 0

	for _, ch := range searchWord {
		prefix += string(ch)

		// Бинарный поиск позиции, где начинается префикс
		i := sort.Search(len(products)-start, func(i int) bool {
			return products[start+i] >= prefix
		}) + start

		matches := []string{}
		for j := i; j < len(products) && len(matches) < 3 && strings.HasPrefix(products[j], prefix); j++ {
			matches = append(matches, products[j])
		}

		res = append(res, matches)
		start = i // оптимизация: сдвигаем начало поиска на следующем шаге
	}

	return res
}
