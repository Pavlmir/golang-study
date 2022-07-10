/*
На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - вывести Right иначе - вывести Wrong

Маленькая подсказка: fmt.Scan() считывает строку до первого пробела, вы можете считать полностью строку за раз с помощью bufio:

text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main1() {
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    asRunes := []rune(text)
    firstLetter := string(asRunes[0])
    if firstLetter == strings.ToUpper(firstLetter) && strings.HasSuffix(text, ".") {
        fmt.Println("Right")
    } else {
         fmt.Println("Wrong")
    }  
}

func main2_5_1() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runes := []rune(text)
	if unicode.IsUpper(runes[0]) && runes[len(runes)-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}

func palindrom() {
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    asRunes := []rune(text)
    j := len(asRunes)
    isPalindrom := false
    for i := 0; i<j; i++ {
        if asRunes[i] != asRunes[j-1] {
            isPalindrom = false
            break
        }  
        isPalindrom = true
        j = j-1
    }
    if isPalindrom {
        fmt.Println("Палиндром")
    } else {
        fmt.Println("Нет")
    }
}

func main2_5_2() {
	var s, r string
	fmt.Scan(&s)
	for _, i := range s {
		r = string(i) + r
	}
	if s == r {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}

func findSubstr() {
	var str, substr string
	fmt.Scan(&str, &substr)
    byteIndex := strings.Index(str, substr)
    if byteIndex == -1 {
        fmt.Println(-1)
    } else {
        index := utf8.RuneCountInString(str[:byteIndex])
        fmt.Println(index)
    }
}

func main2_5_4() {
	var str string
	fmt.Scan(&str)
    asRunes := []rune(str)
    var newWord []rune
    for i := 0; i<len(asRunes); i++ {
        if i % 2 != 0 {
            newWord = append(newWord, asRunes[i])
        }
    }
    fmt.Println(string(newWord))
}

func main2_5_5() {
	var str string
	fmt.Scan(&str)
    asRunes := []rune(str)
    var newWord []rune
    for _, el :=range asRunes {
        if strings.Count(str, string(el)) == 1 {
            newWord = append(newWord, el)
        }
    }
    fmt.Println(string(newWord))
}

func main2_5_6() {
	var str string
	fmt.Scan(&str)
    asRunes := []rune(str)
    alpha := "abcdefghijklmnopqrstuvwxyz"
    var IsLetterNumber = regexp.MustCompile(`.*([a-zA-Z].*[0-9]|[0-9].*[a-zA-Z]).*`).MatchString
    isOk := true
    for _, el :=range asRunes {
        if !(unicode.IsLetter(el) || unicode.IsDigit(el)){
            isOk = false
        }
        if !(strings.Contains(alpha, strings.ToLower(string(el))) || unicode.IsDigit(el)){
            isOk = false
        }
    }
    if utf8.RuneCountInString(str) >= 5 && IsLetterNumber(str) && isOk {
        fmt.Println("Ok")      
    } else {
        fmt.Println("Wrong password")
    }
}

func main2_5_6_better() {
    var str string
	fmt.Scan(&str)
    asRunes := []rune(str)
    isOk := true
    for _, el :=range asRunes {
        if !(unicode.Is(unicode.Latin, el) || unicode.IsDigit(el)){
            isOk = false
            break
        }
    }
    if len(asRunes) >= 5 && isOk {
        fmt.Println("Ok")      
    } else {
        fmt.Println("Wrong password")
    }
}