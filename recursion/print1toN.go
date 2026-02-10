package main

import "fmt"

func main (){
	n := 5
	i:=1
	oneToN(n,i)

	//sum
	fmt.Println(sumN(n))
}
func oneToN (n int, i int) int{
	if i > n {
		return 0
	}
	fmt.Println(i)
	return oneToN(n, i+1)
}

func sumN (n int)int{
	if n == 0 || n == 1{
		return n
	}
	return n + sumN(n-1)
}

