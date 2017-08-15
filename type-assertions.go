package main
import "fmt"

func main(){
  // 型アサーションの例（インターフェースにstring型をアサーション）
  var i interface{} = "hello"
  st := i.(string)
  fmt.Println(st)

  i = 10
  it := i.(int)
  fmt.Println(it)
  st = i.(string)
  fmt.Println(st)
}
