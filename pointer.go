package main
import "fmt"

func main(){
  var pointer *int
  var n int = 100
  pointer = &n
  fmt.Printf("nのアドレス:%v(%T)\n", &n, &n)
  fmt.Printf("pointerの値:%v(%T)\n", pointer, pointer)
  fmt.Printf("nの値:%v(%T)\n", n, n)
  fmt.Printf("pointerの宛先:%v(%T)\n", *pointer, *pointer)
}
