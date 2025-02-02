package main

func isSubsequence(s string, t string) bool {
	lenS := len(s)
	currIdxS := 0
	for i := 0; i < len(t) && currIdxS < lenS; i++ {
		if t[i] == s[currIdxS] {
			currIdxS++
		}
	}
	return currIdxS == len(s)
}

func main() {

}
