package main

import "fmt"

func do(i interface{}){
  switch v:= i.(type){
    case int:
      fmt.Printf("This is int! : %v(%T)\n", v, v)
    case string:
      fmt.Printf("This is string! : %v(%T)\n", v, v)
    default:
      fmt.Printf("Unknown Type.. : %v(%T)\n", v, v)
  }
}

func main(){
  do(10)
  do("hey")
  do(true)
}
