package learn_strings

import (
	bytes2 "bytes"
	"runtime"
	"testing"
)

func TestExampleBytesBuffer(t *testing.T) {
	bytes := bytes2.Buffer{}
	bytes.WriteString("hello")
	bytes.WriteString("world")
	println(bytes.String())
	b := make([]byte, 12)
	read, err := bytes.Read(b)
	println(read)
	if err != nil {
		return
	}
	println(string(b))
	println(bytes.String())
	runtime.Gosched()
}
