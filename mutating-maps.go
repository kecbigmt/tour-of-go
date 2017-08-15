package main
import "fmt"

func main(){
  m := make(map[string]int)
  m["Answer"] = 42
  fmt.Println("The value:", m["Answer"])

  m["Answer"] = 48
  fmt.Println("The value:", m["Answer"])

  delete(m, "Answer")
  fmt.Println("The value:", m["Answer"])

  elem, ok := m["Answer"]
  fmt.Println("The value:", elem, "Present?", ok)

  m["Answer"] = 72
  elem, ok = m["Answer"]
  fmt.Println("The value:", elem, "Present?", ok)
}
