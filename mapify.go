package main

import ("fmt"; "strings"; "unicode")

func main() {

	var text string = "Call me Ishmael. Some years ago- never mind how long precisely- having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world. It is a way I have of driving off the spleen and regulating the circulation. Whenever I find myself growing grim about the mouth; whenever it is a damp, drizzly November in my soul; whenever I find myself involuntarily pausing before coffin warehouses, and bringing up the rear of every funeral I meet; and especially whenever my hypos get such an upper hand of me, that it requires a strong moral principle to prevent me from deliberately stepping into the street, and methodically knocking people's hats off- then, I account it high time to get to sea as soon as I can. This is my substitute for pistol and ball. With a philosophical flourish Cato throws himself upon his sword; I quietly take to the ship. There is nothing surprising in this. If they but knew it, almost all men in their degree, some time or other, cherish very nearly the same feelings towards the ocean with me."

	var hashed map[string]int = NMapify(text, 3)

	fmt.Println(hashed)
}

func Mapify(input string) map[string]int {
	// Accept string input and return map of frequency.

    var (
    	word      string
        frequency = make(map[string]int)
        f = func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
    )

	for _, word = range strings.FieldsFunc(input, f) {
		
		frequency[word]++
		
	}

	return frequency

}

func DiMapify(input string) map[string]int {

	var (
		pair = []string{"", ""}
		frequency = make(map[string]int)
        f = func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
	)

	for index, word := range strings.FieldsFunc(input, f) {
		pair = append([]string{word}, pair[0])
		fmt.Println(pair)
		if index % 2 == 0 {
			frequency[pair[0] + " " + pair[1]]++
		}
	}

	return frequency

}

func TriMapify(input string) map[string]int {

	var (
		pair = []string{"", "", ""}
		frequency = make(map[string]int)
        f = func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
	)

	for index, word := range strings.FieldsFunc(input, f) {
		pair = append([]string{word}, pair[:len(pair)-1]...)
		fmt.Println(pair)
		if index % 3 == 0 {
			frequency[pair[0] + " " + pair[1] + " " + pair[2]]++
		}
	}

	return frequency
	
}

// Turn a string into a map with any number of words as key.
// Then count the instances of those pairs/tuples etc.
func NMapify(input string, words int) map[string]int {

	var (
		firstword = make([]string, 1, 1)
		tuple = make([]string, words, words)
		frequency = make(map[string]int)
        f = func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
	)

	for index, word := range strings.FieldsFunc(input, f) {
		firstword[0] = word
		tuple = append(firstword, tuple[:len(tuple)-1]...)

		if index % 2 == 0 {

			frequency[strings.Join(tuple, " ")]++
		}
	}

	return frequency
}