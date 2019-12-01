package exercise

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/tour/reader"
	"golang.org/x/tour/wc"
)

func sqrt(x float64) float64 {
	z := x / 2
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("On step %v result is %v \n", i, z)
	}
	return z
}

// RunSqrt run exercise with for
func RunSqrt() {
	fmt.Println(sqrt(2))
}

func wordCount(s string) map[string]int {
	a := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range a {
		_, ok := m[v]
		if ok == false {
			m[v] = 1
		} else {
			m[v] = m[v] + 1
		}
	}
	return m
}

// RunWordCount run exercise with map
func RunWordCount() {
	wc.Test(wordCount)
}

func fibonacci() func() int {
	next := 0
	current := 0
	return func() int {
		if next == 0 {
			next = 1
		} else {
			old := next
			next += current
			current = old
		}
		return current
	}
}

// RunFibonacci run exercise with closure
func RunFibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

type ipAddr [4]byte

func (ip ipAddr) String() string {
	return fmt.Sprintf("\"%v.%v.%v.%v\"", ip[0], ip[1], ip[2], ip[3])
}

// RunIPAddr run exercise for Stringer
func RunIPAddr() {
	hosts := map[string]ipAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func sqrt2(x float64) (float64, error) {
	if x < 0 {
		return 0, errNegativeSqrt(x)
	}
	z := x / 2
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("On step %v result is %v \n", i, z)
	}
	return z, nil
}

// RunSqrtError run exercise with error
func RunSqrtError() {
	fmt.Println(sqrt2(2))
	fmt.Println(sqrt2(-2))
}

type myReader struct{}

func (m myReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

// RunReader run exercise with reader
func RunReader() {
	reader.Validate(myReader{})
}

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return
}

// RunRot13 run exercise with wrapper
func RunRot13() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
