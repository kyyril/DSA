package main
import "fmt"

func main (){
	matrix := [][]int {
		{1,2},
		{3,4},
	}
	fmt.Println(det2x2(matrix))

}

func det2x2(m [][]int) int {
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}