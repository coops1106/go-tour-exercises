package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read (b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		if b[i] < 'A' || b[i] > 'z' { //ignore non A-z
			continue
		}
		if b[i] < 'A' + 13  || b[i] > 'Z' && b[i] < 'a' + 13 {
			b[i] += 13
		} else {
			b[i] -= 13
		}
	}
	return len(b), err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

