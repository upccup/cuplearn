package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var sss string
	sss = parseURL("http://127.0.0.1:53153")
	fmt.Println(sss)
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
}

func parseURL(s string) string {
	var sss string
	switch {
	case strings.HasPrefix(s, "ws://"):
		sss = strings.Replace(s, "http", "ws", -1)
	case strings.HasPrefix(s, "wss://"):
		sss = strings.Replace(s, "https", "wss", -1)
	default:
	}
	return sss
}

// Compare returns an integer comparing two strings lexicographocally
// the result will be 0 if a==b  -1 if a <b and +1 if a>b
// Compare in included only symmetry with package bytes it is usually clearer and always
// faster to use the built-in string comparison operators ==, >, < adn so on
func Compare(a, b string) int {
	fmt.Println(strings.Compare("a", "b"))      // -1
	fmt.Println(strings.Compare("ba", "b"))     // 1
	fmt.Println(strings.Compare("aa", "b"))     // -1
	fmt.Println(strings.Compare("a", ""))       // 1
	fmt.Println(strings.Compare("aaaa", "aab")) // -1
	return strings.Compare(a, b)
}

// Contains reports whether substr is within s
func Contains(s, substr string) bool {
	fmt.Println(strings.Contains("seafood", "foo")) // true
	fmt.Println(strings.Contains("seafood", "bar")) // false
	fmt.Println(strings.Contains("seafood", ""))    // true
	fmt.Println(strings.Contains("", ""))           // true
	return strings.Contains(s, substr)
}

// ContainsAny reports whether any Unicode code points in chars within s
func ContainsAny(s, chars string) bool {
	fmt.Println(strings.ContainsAny("team", "i"))        // false
	fmt.Println(strings.ContainsAny("failure", "u & i")) // true
	fmt.Println(strings.ContainsAny("foo", ""))          // false
	fmt.Println(strings.ContainsAny("", ""))             // false
	return strings.ContainsAny(s, chars)
}

// ContainRune reports whether the Unicode code point r is within s
func ContainsRune(s string, r rune) bool {
	fmt.Println(strings.ContainsRune("seafood", 12))    // false
	fmt.Println(strings.ContainsRune("seafood12", 12))  // false
	fmt.Println(strings.ContainsRune("seafood12", 97))  // true 97 ->a
	fmt.Println(strings.ContainsRune("seafood12", 111)) // true  111->o
	return strings.ContainsRune(s, r)
}

// Count counts the number of mon-overlapping instances of sep in s.
// if sep is an empty string, Count return 1+ the number of Unicode points in s
func Count(s, sep string) int {
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("five", ""))    // 5
	fmt.Println(strings.Count("", "e"))       // 0
	return strings.Count(s, sep)
}

// EqulaFold reports whether s and t, interpreted as UTF-8 strings are equal under Unicode case-folding
func EqualFold(s, t string) bool {
	fmt.Println(strings.EqualFold("Go", "go"))    // true
	fmt.Println(strings.EqualFold("abc", "abcd")) // false
	fmt.Println(strings.EqualFold("acc", "ACC"))  // true
	return strings.EqualFold(s, t)
}

// Field splits the string s around each instance of one or more consecutive white space characters,
// as defined by unicode IsSpace, returning an array of substrings of s or an empty list if s contains
// only white space
func Fields(s string) []string {
	fmt.Printf("Fields are: %q", strings.Fields("   foo bar    baz     ")) // Fields are: ["foo" "bar" "baz"]
	return strings.Fields(s)
}

// FieldsFunc splits the string s at each run of Unicode code c satisfying f(c) and returns an array of slices of s.
// if all code points in s satisfy f(c) or the string is empty an empty slice is returned. FieldsFunc makes no guarantees
// about the order in which is calls f(c). if f does not return consistent results for given c FieldsFunc may crash
func FieldsFunc(s string, f func(rune) bool) []string {
	ft := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	fmt.Printf("Field are: %q", strings.FieldsFunc("  foo1;bar2*baz3...", ft)) // Fields are: ["foo1" "bar2" "baz3"]
	return strings.FieldsFunc(s, f)
}
