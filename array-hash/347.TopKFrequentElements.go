package main

import "fmt"

func main (){
	nums := []int {1,1,2,2,3,4}
	k := 2
	topKFrequent(nums, k)
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int] int)
	for _, n := range nums {
		freq[n]++
	}
	fmt.Println(freq)

	// bucket index = frequency
	// bucket value = list of numbers
	buckets := make([][]int, len(nums)+1)
    for num, f := range freq {
        buckets[f] = append(buckets[f], num)
    }
	fmt.Println(buckets)

	res:= make([]int, 0 , k)
	for i := len(buckets) - 1; i >= 0 && len(res) < k; i-- {
        for _, num := range buckets[i] { //from back idx
            if len(res) < k {
                res = append(res, num)
            } else {
                break
            }
        }
    }
	fmt.Println(res)
	return res 
}