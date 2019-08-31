package main
import (
  "fmt"
  )

// 重複した要素を削除して返却
/*
func removeDuplicate1(args []string) []string {
    results := make([]string, 0, len(args))
    encountered := map[string]bool{}
    for i := 0; i < len(args); i++ {
      fmt.Println(encountered)
        if !encountered[args[i]] {
            encountered[args[i]] = true
            results = append(results, args[i])
        }
    }
    return results
}
*/

// 差集合（set1-set2）
func MakeIntersection(set1 []string, set2 []string) []string{
  dict := map[string]bool{}
  for _, v := range set1{
    dict[v] = true
  }
  for _, v := range set2{
    if dict[v] {
      delete(dict, v)
    }
  }
  result := make([]string, len(dict), len(dict))
  i := 0
  for k, _ := range dict {
    fmt.Println(i, k)
    result[i] = k
    i++
  }
  return result
}

func main() {
  s1 := []string{"a", "b", "c", "c"}
  s2 := []string{"c", "d", "e"}
  i := MakeIntersection(s1, s2)
  fmt.Println(i, len(i))
}
