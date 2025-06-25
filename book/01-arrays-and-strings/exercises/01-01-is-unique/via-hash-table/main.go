// ENG:
// 		Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
// 		cannot use additional data structures?
// 		Hints: #44, #117, #132
// RU:
// 		Реализуйте алгоритм, определяющий, все ли символы в строке встречаются только один раз.
// 		А если при этом запрещено использование дополнительных структур данных?

package main

// time complexity  = O(n)
// space complexity = O(n)
func isUnique(s string) bool {
	// 0. corner cases
	if s == "" {
		return true
	}

	// 1. create hash table for unique symbols
	uniqueSymbols := map[rune]bool{}

	// 2. loop over chars in string
	for _, r := range s {
		// 3. return false if there are any duplicates
		if uniqueSymbols[r] {
			return false
		}
		// 4. add unique char to hash table
		uniqueSymbols[r] = true
	}
	// 5. if we reach this line that means there are no duplicates
	return true
}

func main() {
	// test - TRUE
	println("5 TRUE")
	println(isUnique("abc"))
	println(isUnique("123qwerty456"))
	println(isUnique(" ")) // space
	println(isUnique("	")) // tab
	println(isUnique(""))

	// test - FALSE
	println("4 FALSE")
	println(isUnique("abca"))
	println(isUnique("12345678901"))
	println(isUnique("  ")) // 2 spaces
	println(isUnique("		")) // 2 tabs
}
