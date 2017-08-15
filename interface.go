package main
import "fmt"

type Animal interface{
  Bark()
}

type Dog struct{}
func (v Dog) Bark(){
  fmt.Println("わんわん")
}

type Cat struct{}
func (v Cat) Bark(){
  fmt.Println("にゃあ")
}

func MakeAnimalBark(a Animal){
  fmt.Println("鳴け！")
  a.Bark()
}

func main(){
  dog := new(Dog)
  cat := new(Cat)
  MakeAnimalBark(dog)
  MakeAnimalBark(cat)
}
