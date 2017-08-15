package main
import "fmt"

func main(){
  names := [4]string{
    "John",
    "Paul",
    "George",
    "Ringo",
  }
  fmt.Printf("names:%v(%T)\n", names, names)

  a := names[0:2]
  b := names[1:3]
  fmt.Printf("a:%v(%T)\n", a, a)
  fmt.Printf("b:%v(%T)\n", b, a)

  b[0] = "XXX"
  fmt.Println("--b[0] changed--")

  fmt.Printf("names:%v(%T)\n", names, names)
  fmt.Printf("a:%v(%T)\n", a ,a)
  fmt.Printf("b:%v(%T)\n", b, b)
}
