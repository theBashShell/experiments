package main

import "fmt"

func catchpanic() {
	r := recover()
	if r != nil{
	fmt.Println(r)
	}
}

func factorizor(factor, num int) func() int {
	return func() int {
		return num * factor
	}
}

func p(a ...interface{}) {
	for _, val := range a {
		fmt.Printf("**%v\n", val)
	}
}

func main() {
	defer catchpanic()
	slc := []int{}
	slc = append(slc, 20, 303)
	fmt.Println(slc, len(slc))
	// slc[0] = 0
	println("wow", slc)
	if slc[0] == 0 {panic("noooo, not zero")}

	f := factorizor(10, 8)
	fmt.Println(f())
	p(20, "hello", '5', slc)
}