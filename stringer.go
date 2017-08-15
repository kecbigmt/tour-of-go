package main
import "fmt"

type Person struct{
  Name string
  Age int
}
func (p Person) String() string{
  return fmt.Sprintf("%v(%v years)", p.Name, p.Age)
}

func main(){
  fmt.Println(Person{"Michael", 43})
  fmt.Println(Person{"Scalet", 32})
}
