package main

import (
  "io"
  "os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}
func (r13 rot13Reader) Read(b []byte) (int, error){
  r := r13.r
  n, err := r.Read(b)
  for i, v := range b[:n]{
    switch {
    case (v>='A' && v<='M')||(v>='a' && v<='m'):
      b[i] = v + uint8(13)
    case (v>='N' && v<='Z')||(v>='n' && v<='z'):
      b[i] = v - uint8(13)
    }
  }
  return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
