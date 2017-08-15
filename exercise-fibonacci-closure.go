package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n:=0
	f0:=0
	f1:=1
	return func() int{
		switch n{
			case 0:
				n += 1
				return 0
			case 1:
				n += 1
				return 1
			default:
				n += 1
				f2 := f0+f1
				f0 = f1
				f1 = f2
				return f2
		}

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
