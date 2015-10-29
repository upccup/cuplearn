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

// HasPrefix tests whether the string s begins with prefix
func HasPrefix(s, prefix string) bool {
	fmt.Println(strings.HasPrefix("*seafood", "*sea")) // ture
	fmt.Println(strings.HasPrefix("saafood", "bar"))   // false
	fmt.Println(strings.HasPrefix("seafood", ""))      // true
	fmt.Println(strings.HasPrefix("", ""))             // true
	return strings.HasPrefix(s, prefix)
}

// HasSuffix tests whether the string s ends with suffix
func HasSuffix(s, suffix string) bool {
	fmt.Println(strings.HasSuffix("*seafood", "food")) // true
	fmt.Println(strings.HasSuffix("seafood", "bar"))   // false
	fmt.Println(strings.HasSuffix("seafood", ""))      // true
	fmt.Println(strings.HasSuffix("", ""))             // true
	return strings.HasSuffix(s, suffix)
}

// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s
func Index(s, sep string) int {
	fmt.Println(strings.Index("chicken", "ken"))       // 4
	fmt.Println(strings.Index("chickenkenken", "ken")) // 4
	fmt.Println(strings.Index("chicken", "dmr"))       // -1
	fmt.Println(strings.Index("chicken", ""))          // 0
	return strings.Index(s, sep)
}

// IndexAny returns the index of the first instance of any Unicode code point from chars in s
// or -1 if no Unicode code point from chars is present in s
func IndexAny(s, chars string) int {
	fmt.Println(strings.IndexAny("chicken", "aeiouy"))    // 2
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))      // -1
	fmt.Println(strings.IndexAny("crwth", "axrrrrrrrr"))  // 1
	fmt.Println(strings.IndexAny("crwth", "acsxxxrrrrr")) // 0
	return strings.IndexAny(s, chars)
}

// IndexByte returns the index of thr first instance of c in s
// or -1 if c is not present in s
func IndexByte(s string, c byte) int {
	fmt.Println(strings.IndexByte("crwth", 99))  // 0
	fmt.Println(strings.IndexByte("crwth", 114)) // 1
	fmt.Println(strings.IndexByte("crwth", 67))  // -1
	return strings.IndexByte(s, c)
}

// IndexFunc returns the index into s of the first Unicode code point satisfying f(c)
// or -1 if none do
func IndexFunc(s string, f func(rune) bool) int {
	function := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}

	fmt.Println(strings.IndexFunc("Hello, 世界", function))    // 7
	fmt.Println(strings.IndexFunc("Hello, world", function)) // -1
	return strings.IndexFunc(s, f)
}

// IndexRune returns the index of thr first instance of the Unicode code point r
// or -1 if rune is not present in s
func IndexRune(s string, r rune) int {
	fmt.Println(strings.IndexRune("chicken", 'c')) // 0
	fmt.Println(strings.IndexRune("chicken", 99))  // 0
	fmt.Println(strings.IndexRune("chicken", 'd')) // -1
	return strings.IndexRune(s, r)
}

// Join concatenates the elements of a to create a single string
// The separator string sep is placed between elements in the resulting string
func Join(a []string, sep string) string {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) // foo, bar, baz
	return strings.Join(a, sep)
}

// LastIndex returns the index of the last instance of sep in s
// or -1 if sep is not present in s
func LastIndex(s, sep string) int {
	fmt.Println(strings.Index("go gopher", "go"))         // 0
	fmt.Println(strings.LastIndex("go gopher", "go"))     // 3
	fmt.Println(strings.LastIndex("go gopher", "rodent")) // -1
	return strings.LastIndex(s, sep)
}

// LastIndexAny returns the index of the last instance of any Unicode code
// point from chars in s, or -1 if no Unicode code point from chars is present in s
func LastIndexAny(s, chars string) int {
	fmt.Println(strings.LastIndexAny("gogopher", "go"))     // 3
	fmt.Println(strings.LastIndexAny("gogopher", "ogh"))    // 5
	fmt.Println(strings.LastIndexAny("gogopher", "gr"))     // 7
	fmt.Println(strings.LastIndexAny("gogopher", "rodent")) // 7
	return strings.LastIndexAny(s, chars)
}

// LastIndexByte returns the index of the last instance of c in s
// or -f if c is not present in s
func LastIndexByte(s string, c byte) int {
	fmt.Println(strings.LastIndexByte("gogopher", 'o')) // 3
	fmt.Println(strings.LastIndexByte("gogopher", 111)) // 3
	fmt.Println(strings.LastIndexByte("gogopher", 112)) // 4
	fmt.Println(strings.LastIndexByte("gogopher", 113)) // -1
	return strings.LastIndexByte(s, c)
}

// LastIndexFunc returns the index into s of the last Unicode code point
// satisfying f(c) or -1 if none do
func LastIndexFunc(s string, f func(rune) bool) int {
	function := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}

	fmt.Println(strings.LastIndexFunc("hello 世界", function))    // 10 一个汉字貌似占3个位置
	fmt.Println(strings.LastIndexFunc("hello world", function)) // -1
	return strings.LastIndexFunc(s, f)
}

// Map returns a copy of the the string s with all its characters modified according to the mapping
// function if mapping returns a negative value, the character is dropped from the string with no replacement
func Map(mapping func(rune) rune, s string) string {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r < 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}

	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher...")) // 'Gjnf oevyyvt naq fyvgul tbcure...
	return strings.Map(mapping, s)
}

// Repeats returns a new string consisting of count copies of the string s
func Repeat(s string, count int) string {
	fmt.Println("ba " + strings.Repeat("na", 2)) // ba nana
	return strings.Repeat(s, count)
}

// Replace returns a copy of the string s with the first n non-overlapping instances
// of old replaced by new. if old is empty it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string
// if n < 0 there is not limit on the number of replacements
func Replace(s, old, new string, n int) string {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      //oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // moo moo moo
	fmt.Println(strings.Replace("oink oink oink", "", "ky", 2))       // kyokyink oink oink
	fmt.Println(strings.Replace("oink oink onik", "", "ky", -1))      // kyokyikynkykky kyokyikynkykky kyokyikynkykky
	return strings.Replace(s, old, new, n)
}

// Split slices s into all substrings separated by sep and returns a slice 	of the substrings
// between those serarators . if sep is enpty Split splits after each UTF-8 sequence . It is
// equivalent to SplitN with a count of -1
func Split(s, sep string) []string {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        // ["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // ["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz", ""))                          // [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))            // [""]
	return strings.Split(s, sep)
}

// SplitN slices s into substrings separated by sep and returns a slices of the substrings
// between those separators. if sep is empty, SplitN splits after each UTF-8 sequence.
// The count determines the bunber of substrings to return
func SplitN(s, sep string, n int) []string {
	fmt.Println("%q\n", strings.SplitN("a,b,c", ",", 2)) // ["a" "b,c"]
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Println("%q (nil = %v)", z, z == nil) // [] (nil = true)
	return strings.SplitN(s, sep, n)
}

// SplitAfter slices s into all substrings after each instance of sep and returns
// a slice of those substrings. if sep is empty, SplitAfter splits after each UTF-8 sequence.
// it is equivalents to SplitAfterN with a count of -1
func SplitAfter(s, sep string) []string {
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ",")) // ["a," "b," "c"]
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ""))  // ["a" "," "b" "," "c"]
	return strings.SplitAfter(s, sep)
}

// SplitAfterN slices s into substrings after each instance of sep and returns a slice of those substrings
// if sep is enpty, SplitAfterN splits after each UTF-8 sequence The count deterrmines the number of
// substrings to return
// n > 0 : at most n substrings; the last substring will be the unsplit remainder
// n = 0 : the result is nil (zero substrings)
// n < 0 : all substrings
func SplitAfterN(s, sep string, n int) []string {
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 1)) // ["a,b,c"]
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2)) // ["a," "b,c"]
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 3)) // ["a," "b," "c"]
	return strings.SplitAfterN(s, sep, n)
}

// Title returns a copy of the string s with all Unicode letters that begin words mapped to their title case
// BUG(rsc) The rule Title uses for word boundaries does not handle Unicode puncyuation properly
func Title(s string) string {
	fmt.Println(strings.Title("her royal highness")) // Her royal Highness
	return strings.Title(s)
}

// ToLower returns a copy of the string s with all Unicode letters mapped to their lower case
func ToLower(s string) string {
	fmt.Println(strings.ToLower("GoPher")) // gopher
	return strings.ToLower(s)
}

// ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their lower case
// giving priority to the special casing rules
func ToLowerSpecial(_case unicode.SpecialCase, s string) string {
	fmt.Println(strings.ToLowerSpecial(unicode.AzeriCase, "ŞĞÜÖIİ")) // şğüöıi
	return strings.ToLowerSpecial(_case, s)
}

// ToTitle returns a copy of the string s with all Unicode letters mapped to their title case
func ToTitle(s string) string {
	fmt.Println(strings.ToTitle("loud noises")) // LOUD NOISES
	fmt.Println(strings.ToTitle("хлеб"))        // ХЛЕБ
	return strings.ToTitle(s)
}

// ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their title case
// giving priority to the special casing rules
func ToTitleSpecial(_case unicode.SpecialCase, s string) string {
	fmt.Println(strings.ToTitleSpecial(unicode.AzeriCase, "şğüöıi")) // ŞĞÜÖIİ
	return strings.ToTitleSpecial(_case, s)
}

// ToUpper returns a copy of the string s with all Unicode letters mapped to their upper case
func ToUpper(s string) string {
	fmt.Println(strings.ToUpper("GoPher")) // GOPHER
	return strings.ToUpper(s)
}

//ToUpperSpecial returns a copy of thr string s with all Unicode letters mapped to their upper case
// giving priority to the special casing rules
func ToUpperSpecial(_case unicode.SpecialCase, s string) string {
	fmt.Println(strings.ToUpperSpecial(unicode.AzeriCase, "şğüöıi")) // ŞĞÜÖIİ
	return strings.ToUpperSpecial(_case, s)
}

// Trim returns a slice of the string s with all leading and trailing Unicode code points
// contained in cutest removed
func Trim(s string, cutest string) string {
	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))     // ["Achtung! Achtung"]
	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung! Achtung! !!! @@@ ", "!@")) // [" !!! Achtung! Achtung! !!! @@@ "]
	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung! Achtung! !!! ", ""))       // [" !!! Achtung! Achtung! !!! "]
	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung! Achtung! !!! ", " "))      // ["!!! Achtung! Achtung! !!!"]
	return strings.Trim(s, cutest)
}

// TrimFunc returns a slice of the string s with all leading adn trailing Unicode code point
// c satisfying f(c) removed
func TrimFunc(s string, f func(rune) bool) string {
	inner_func := func(r rune) bool {
		return r <= 'c' || r >= 'i'
	}
	fmt.Println(strings.TrimFunc("abcdefghijk", inner_func)) // defgh
	return strings.TrimFunc(s, f)
}
