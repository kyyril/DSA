package main 
import "fmt"

func main () {
	report := make(map[string] int)
	report["wahyu"] = 90
	report["asep"] = 95
	report["jamal"] = 100
	fmt.Println(report)
	
	delete(report, "wahyu")
	if _, ok := report["wahyu"]; ok {
		fmt.Println("wahyu is here")
		}else{
			fmt.Println("wahyu not found")
	}
	fmt.Println(report)
}