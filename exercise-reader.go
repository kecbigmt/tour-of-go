package main
import (
	"../golang.org/x/tour/reader"
)
//"io"
//"strings"

/*type MyReader struct{
	s string
	i int64
}
func (m *MyReader) Read (b []byte) (n int, err error){
	if m.i >= int64(len(m.s)){
		return 0, io.EOF
	}
	n = copy(b, m.s[m.i:])
	m.i += int64(n)
	return
}*/
type MyReader struct{}
func (m MyReader) Read (b []byte) (n int, err error){
	b[0] = 'A'
	n = 1
	return
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	/*var s string
	for {
		s = strings.Join([]string{"A","A"},"")
	}
	r := &MyReader{s,0}
	reader.Validate(r)*/
	reader.Validate(MyReader{})
}
