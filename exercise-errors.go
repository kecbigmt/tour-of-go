package main
import (
  "fmt"
  "math"
  "errors"
)

/*
type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string{
  return fmt.Sprintf("cannot Squrt negative number:%v", float64(e))
}*/

func Sqrt(x float64) (float64, error) {
  if x<0{
    return 0, errors.New(fmt.Sprintf("cannot Squrt negative number:%v", float64(x)))
  }else{
    var zn float64 = 1.0
    var zn1 float64
      for {
        zn1 = zn-(zn*zn-x)/2.0*zn
        if math.Abs(zn-zn1) < 0.01{
          break
        }
        zn = zn1
      }
    return zn1, nil
  }
}

func main() {
	fmt.Println(Sqrt(4.0))
	fmt.Println(Sqrt(-2.0))
}
