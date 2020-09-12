package main

import (
	"fmt"
	"io"
	"os"
)

type alphaReader struct {
	src string
	cur int
}

func newAlphaReader(src string) *alphaReader {
	return &alphaReader{src: src}
}

type alphaReader1 struct {
	src io.Reader
}

func newReader(src io.Reader) *alphaReader1 {
	return &alphaReader1{src}
}
func alpha(r byte) byte {
	if (r >= 'A' && r < 'Z') || (r >= 'a' && r < 'z') {
		return r
	}
	return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {

	if a.cur >= len(a.src) {
		return 0, io.EOF
	}

	x := len(a.src) - a.cur

	n, bound := 0, 0

	if x >= len(p) {
		bound = len(p)
	} else {
		bound = x
	}

	buf := make([]byte, bound)

	for n < bound {
		if char := alpha(a.src[a.cur]); char != 0 {
			buf[n] = char
		}
		n++
		a.cur++
	}

	copy(p, buf)
	return n, nil
}

func (a *alphaReader1) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	count, err := a.src.Read(p) // p has now source data
	if err != nil {
		return count, err
	}
	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] <= 'Z') ||
			(p[i] >= 'a' && p[i] <= 'z') {
			continue
		} else {
			p[i] = 0
		}
	}
	return count, io.EOF
}
func customIoReader() {
	reader := newAlphaReader("Hello! It's 9am, where is the sun?")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()

	file, _ := os.Open("./CustomIOReader.go")
	alpha := newReader(file)
	io.Copy(os.Stdout, alpha)
	fmt.Println()

}
