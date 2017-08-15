package main
import "fmt"

func Pic(dx, dy int) [][]uint8 {
    a := make([][]uint8, dy)
	for i := range a{
	    //[i][]uint8 = make([]uint8, dx)
      a[i] = make([]uint8, dx)
	}
  for i := range a{
    for j:= range a[i]{
      a[i][j] = uint8((i+j)/2)
    }
  }
  return a
}

func main(){
  fmt.Println(Pic(10,10))
}
