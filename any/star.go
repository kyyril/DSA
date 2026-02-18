package main

import "fmt"

func main() {
	n := 5
  fiveXfive(n) //5x5
  fmt.Println("==================")
  triangle(n)
}
func fiveXfive(n int) {
  for i := 1; i <= n; i++ {
    for s:=1; s <= n; s++ {
      fmt.Print("# ")
    }
		fmt.Println()
	}
}

func triangle(n int) {
  for i:=1; i <= n; i++ {
    for s:=1; s <= n-i; s++ {
      fmt.Print(" ")
    }
    for j:=1; j <= 2*i-1; j++ {
      fmt.Print("*")
    }
    fmt.Println()
  }
}