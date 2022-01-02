package main

import (
	"fmt"
	"constraints"
)

type Num int 

type Ordered interface{
	~int | ~int8 | ~int16 
}

func (n *Num)String() string {
	//var placer any = *n

	var i int = int(*n)
	return fmt.Sprintf("%d", i)
}

func main() {
	Print(3)
	number := Num(1024)
	//name := "angel"
	vals := []fmt.Stringer { &number}
	text := Stringy(vals)
	fmt.Println(text)
	fmt.Println(Max(-1000, 32, 88, 1000, 32, 234))
}

func Print[T any](value T){
	fmt.Println(value)
}

func Stringy[T fmt.Stringer] (s []T) (ret []string){
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

func Max[E constraints.Ordered] (input ...E) (max E){
	if len(input) < 1 {
		return
	}

	max = input[0]

	for _, val := range input {
		if val >= max {
			max = val 
		}
	}
	return
}
