package main
import (
  "fmt"
  "sync"
  "time"
)

type SafeCounter struct{
  v map[string]int
  mux sync.Mutex
}

func (c *SafeCounter) Inc(key string){
  c.mux.Lock()
  c.v[key]++
  c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int{
  c.mux.Lock()
  defer c.mux.Unlock()
  return c.v[key]
}

func main(){
  c := SafeCounter{v: make(map[string]int)}
  for i := 0; i <1000; i++ {
    go c.Inc("some key")
    go Hello()
  }
  // 複数のgoroutineが走っている途中でも変数に安全にアクセスできる
  time.Sleep(time.Nanosecond)
  fmt.Println("1 nanosecond elapsed. counter=", c.Value("some key"))
  time.Sleep(time.Nanosecond)
  fmt.Println("2 nanoseconds elapsed. counter=", c.Value("some key"))
}
