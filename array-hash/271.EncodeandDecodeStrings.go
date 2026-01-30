// Length-Prefix Encoding
// Used in:
// network protocols
// databases
// message queues
// real backend systems

package main

import (
	"fmt"
	"strconv"
)

func main() {
	codec := Constructor()
	encoded := codec.Encode([]string{"hello", "world","team"})
	decoded := codec.Decode(encoded)
	fmt.Println(encoded)
	fmt.Println(decoded)
}
type Codec struct {
}

func Constructor () Codec{
	return Codec{}
}

func (c *Codec) Encode(strs []string) string {
	var res []byte
	for _, s := range strs {
		res = append(res, []byte(strconv.Itoa(len(s)))...)
		res = append(res, '#')
		res = append(res, []byte(s)...)
	}
	return string(res)
}


// 5#hello5#world
// use pointer
func (c *Codec) Decode(str string) string {
	var res string
	i := 0
	for i < len(str) {
		j := i
		for str[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(str[i:j])
		j++
		res += str[j:j+length]

		i = j + length //move next word
		if i < len(str) {
			res += " "
		}
	}
	return res
}