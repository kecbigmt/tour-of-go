package main
import "fmt"

func Do(i interface{}) {
  switch v:=i.(type) {
  case int:
    fmt.Printf("INT:%b(%T)\n", v, v)
  case string:
    fmt.Printf("STRING:%v(%T)\n", v, v)
  default:
    fmt.Printf("DEFAULT:%v(%T)\n", v, v)
  }
}

func main(){
  Do(100)
}
