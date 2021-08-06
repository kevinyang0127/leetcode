package main

import "fmt"

func main() {
	A := []int{3, 5, 1, 2, 3}
	B := []int{3, 6, 3, 3, 4}
	fmt.Println(minDominoRotations(A, B))
}

func minDominoRotations(A []int, B []int) int {
	var recodeANumCount [6]int //紀錄A array 每種數字出現次數
	var recodeBNumCount [6]int //紀錄B array 每種數字出現次數
	var eachNumCount [6]int    //紀錄A+B array每種數字出現次數，若同行數字一樣只加一次，
	//最後若交換可行，那必定那個可行的數字數量會等於input array長度
	length := len(A)
	for i := 0; i < length; i++ {
		recodeANumCount[A[i]-1]++
		recodeBNumCount[B[i]-1]++
		eachNumCount[A[i]-1]++
		if A[i] != B[i] {
			eachNumCount[B[i]-1]++
		}
	}

	for i := 0; i < 6; i++ {
		if eachNumCount[i] == length {
			//若交換可行，最小交換次數為 array長度 - 哪個array擁有最多該數字的數量
			if recodeANumCount[i] > recodeBNumCount[i] {
				return length - recodeANumCount[i]
			}

			return length - recodeBNumCount[i]
		}
	}
	return -1
}
