package main

import "fmt"

type void struct{}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	s2 := "AAAAAAAAAAAAA"
	s3 := "GATGGATACTGCATTCGAACCAGAGCCGGCTTTTGCGGGACTAGCATGAGGGACTTGGCTGTCCATCTTTAAATACGAAACCCAA"
	fmt.Println(findRepeatedDnaSequences(s))
	fmt.Println(findRepeatedDnaSequences(s2))
	fmt.Println(findRepeatedDnaSequences(s3))
}

func contains(s []string, str string) bool {
	// not use, for future
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func findRepeatedDnaSequences(s string) []string {
	var result []string
	var myMap = makeAllSubStrToMap(s)

	for key, value := range myMap {
		if value > 1 {
			result = append(result, key)
		}
	}

	return result
}

func makeAllSubStrToMap(s string) map[string]int {
	myMap := make(map[string]int)
	sLen := len(s)

	for i := 0; i < sLen; i++ {
		if i+9 >= sLen {
			break
		}
		nowSubStr := s[i : i+10]
		myMap[nowSubStr]++
	}

	return myMap
}
