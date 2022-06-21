package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if (b >= 'A' && b <= 'M') || (b >= 'a' && b <= 'm') {
		b += 13
	} else if (b >= 'N' && b <= 'Z') || (b >= 'n' && b <= 'z') {
		b -= 13
	}
	return b
}

func (reader rot13Reader) Read(buffer []byte) (int, error) {
	n, e := reader.r.Read(buffer)
	for i := 0; i < n; i++ {
		buffer[i] = rot13(buffer[i])
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
