package main

import "fmt"

func main() {
	fmt.Println(buddyStrings("ab", "ab"))           //false
	fmt.Println(buddyStrings("baaaa", "aaaab"))     //true
	fmt.Println(buddyStrings("baasdaa", "baasdaa")) //true
	fmt.Println(buddyStrings("adxzcz", "uilknjm"))  //false
	fmt.Println(buddyStrings("adxzcz", "adxzzc"))   //true
	fmt.Println(buddyStrings("", "aaaab"))          //false
	fmt.Println(buddyStrings("papapap", "paappap")) //true
	fmt.Println(buddyStrings("abcd", "badc"))       //false
}

func buddyStrings(A string, B string) bool {
	aLen := len(A)

	if aLen != len(B) {
		return false
	}

	notSameTimes := 0
	hasSameWord := false
	var wordArr [27]bool
	var record [3]string
	for i := 0; i < aLen; i++ {
		if wordArr[A[i]-'a'] {
			hasSameWord = true
		} else {
			wordArr[A[i]-'a'] = true
		}

		if A[i] != B[i] {
			record[notSameTimes] = string([]byte{A[i], B[i]})
			notSameTimes++
			if notSameTimes > 2 {
				return false
			}
		}
	}

	if notSameTimes == 2 {
		return record[0][0] == record[1][1] && record[0][1] == record[1][0]
	}

	return notSameTimes == 0 && hasSameWord
}
