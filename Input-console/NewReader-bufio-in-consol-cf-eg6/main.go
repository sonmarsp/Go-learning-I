package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
Пример с сайта но хер пойми как он работает
*/

func main() {
	io := newFastIO(os.Stdin, os.Stdout)
	defer io.Flush()

	for t := io.nextInt(); t > 0; t-- {

	}

	fmt.Println(io)

}

type fastIO struct {
	*bufio.Reader
	*bufio.Writer
	p [1]byte
}

func newFastIO(r io.Reader, w io.Writer) *fastIO {
	return &fastIO{Reader: bufio.NewReader(r), Writer: bufio.NewWriter(w)}
}

func (f *fastIO) nextInt() int {
	f.Read(f.p[:])
	for f.p[0] <= ' ' {
		f.Read(f.p[:])
	}
	neg := false
	if f.p[0] == '-' {
		neg = true
		f.Read(f.p[:])
	}
	n := 0
	for '0' <= f.p[0] && f.p[0] <= '9' {
		n = n*10 + int(f.p[0]) - '0'
		f.Read(f.p[:])
	}
	if neg {
		return -n
	}
	return n
}

func (f *fastIO) nextInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = f.nextInt()
	}
	return a
}

func (f *fastIO) next() string {
	f.Read(f.p[:])
	for f.p[0] <= ' ' {
		f.Read(f.p[:])
	}
	var s []byte
	for f.p[0] > ' ' {
		s = append(s, f.p[0])
		f.Read(f.p[:])
	}
	return string(s)
}

func (f *fastIO) println(a ...interface{}) {
	fmt.Fprintln(f.Writer, a...)
}

func (f *fastIO) print(a ...interface{}) {
	fmt.Fprint(f.Writer, a...)
}

func (f *fastIO) printInts(a ...int) {
	if len(a) > 0 {
		f.print(a[0])
		for _, x := range a[1:] {
			f.print(" ", x)
		}
	}
}

func (f *fastIO) printlnInts(a ...int) {
	f.printInts(a...)
	f.println()
}
