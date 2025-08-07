package main 

import (
	"fmt"
	"os"
	"hello"
	"strings"
	"bufio"
)

func main() {

	if len(os.Args) > 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)

		fmt.Println(t)
	}
	// go run . evelyn rosa < texts.txt
	// test change

	fmt.Println(hello.Say(os.Args[1:]))
	
	// the lenght of a string is the number of bytes required to represent the unicode characters
	// string are immutable
	s := "你好"
	fmt.Printf("%8T %[1]v %d\n", s, len(s))
	fmt.Printf("%8T %[1]v\n", []rune(s))
	
	b := []byte(s)
	fmt.Printf("%8T %[1]v %d\n", b, len(b))

	var sum float64
	var n int

	for {
		var val float64
		if _, err := fmt.Fscanln(os.Stdin, &val); err != nil {
			break
		}
		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	fmt.Println("The average is", sum/float64(n))
	// cd cmd
	// go run . < numbers.txt OR cat numbers.txt | go run .
}