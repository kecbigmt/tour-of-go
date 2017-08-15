package main
import "fmt"

type Vertex struct{
  X int
  Y int
}
var (
  v1 = Vertex{1,2} //has type Vertex
  v2 = Vertex{X:1} // Y:0 is implicit
  v3 = Vertex{} // X:0 and Y:0
  v4 = &Vertex{3,4} // has type *Vertex
)

func main(){
  fmt.Printf("v1:%v(%T)\n", v1, v1)
  fmt.Printf("v1:%v(%T)\n", v2, v2)
  fmt.Printf("v1:%v(%T)\n", v3, v3)
  fmt.Printf("v1:%v(%T)\n", v4, v4)
}
