// ENG:
// 		Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
// 		cannot use additional data structures?
// 		Hints: #44, #117, #132
// RU:
// 		Реализуйте алгоритм, определяющий, все ли символы в строке встречаются только один раз.
// 		А если при этом запрещено использование дополнительных структур данных?

package main

func isUnique(s string) bool {
	// TODO: implement after 5th chapter
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
