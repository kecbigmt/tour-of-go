package main
import (
  "fmt"
  "strings"
)

func WordCount(s string) map[string]int{
  m := map[string]int{}
  f := strings.Fields(s)
  for _, v := range f{
    e, ok := m[v]
    if ok {
      m[v] = e+1
    } else {
      m[v] = 1
    }
  }
  return m
}

func main(){
  fmt.Println(WordCount("I am fine Thank you And you ?"))
}
