package main

import "fmt"

func main (){
	A := [][]int{
		{1, 2},
		{3, 4},
	}

	B := [][]int{
		{5, 6},
		{7, 8},
	}

	addMatrix(A, B)
}
func addMatrix (A, B [][]int) [][]int {
	rows := len(A)
	cols := len(A[0])

	C := make([][]int, rows)
	for i := range C {
		C[i] = make([]int, cols)
	}

	for i:=0; i < rows; i++ {
		for j:=0; j < cols; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	fmt.Println(C)
	return C
}