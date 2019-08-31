package main
import (
  "golang.org/x/tour/tree"
  "fmt"
  )

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  ch <- t.Value
  if t.Left != nil {
    go Walk(t.Left, ch)
  }
  if t.Right != nil {
    go Walk(t.Right, ch)
  }
  return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int)
  ch2 := make(chan int)
  go Walk(t1, ch1)
  go Walk(t2, ch2)
  ar1 := [10]int{}
  ar2 := [10]int{}
  for i := 0; i < 10; i++{
    ar1[i] = <-ch1
    ar2[i] = <-ch2
  }
  count := 0
  for _, v1 := range ar1 {
    for _, v2 := range ar2 {
      if v1 == v2 {
        count += 1
        continue
      }
    }
  }
  if count == 10 {
    return true
  } else{
    return false
  }
}

func main() {
  fmt.Println(Same(tree.New(1), tree.New(1)))
}
