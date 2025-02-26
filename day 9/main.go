package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//converting data type

	var str = "124"
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Println(num) // 124
	}

	var num1 = 124
	var str1 = strconv.Itoa(num1)

	fmt.Println(str1) // "124"

	//method on strings

	//checking string contain
	var isExists = strings.Contains("john wick", "wick")
	fmt.Println(isExists)

	//checking prefix of the text
	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1) // true

	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2) // false

	//checking suffix of the text
	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1) // false

	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2) // true

	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"

	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1) // output: ["the", "dark", "knight"]

	var string2 = strings.Split("batman", "")
	fmt.Println(string2) // output: ["b", "a", "t", "m", "a", "n"]

	var data = []string{"banana", "papaya", "tomato"}
	var str2 = strings.Join(data, "-")
	fmt.Println(str2) // "banana-papaya-tomato"

	
}
