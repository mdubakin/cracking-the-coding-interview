// ENG:
// 		Check Permutation: Given two strings, write a method to decide if one is a permutation of the other.
// 		Hints: #1, #84, #122, #131
// RU:
// 		Для двух строк напишите метод, определяющий, является ли одна строка перестановкой другой.

package main

// time complexity  = O(n)
// space complexity = O(n)
func isPermutation(s1, s2 string) bool {
	// 0. corner cases
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}

	s1Chars := map[rune]int{}
	for _, r1 := range s1 {
		// O(n) via
		// strings.Contains(s2, string(r))

		// 1. fill map with chars
		s1Chars[r1]++
	}

	for _, r2 := range s2 {
		// 2. return false if s2 char doesn't exist in s1
		if _, ok := s1Chars[r2]; !ok {
			return false
		}

		// 3. decrement char count in map
		s1Chars[r2]--
		// 4. if there are no chars left, then delete that key
		if s1Chars[r2] == 0 {
			delete(s1Chars, r2)
		}
	}

	// 5. if s1Chars map is empty, then s1 and s2 are permutations
	if len(s1Chars) == 0 {
		return true
	}
	return false
}

func main() {
	// test - TRUE
	println(isPermutation("abc", "abc"))
	println(isPermutation("abc", "cba"))

	// test - FALSE
	println(isPermutation("abc", "abca"))
	println(isPermutation("abcb", "abca"))
}
