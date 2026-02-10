package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){

	println(len(os.Args))

	if len(os.Args) < 3 {
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

	// go run ./cmd/2-replacer name1 name2 < ./cmd/texts.txt
	
	// the lenght of a string is the number of bytes required to represent the unicode characters
	// strings are immutable
	s := "你好"
	fmt.Printf("%8T %[1]v %d\n", s, len(s))
	fmt.Printf("%8T %[1]v\n", []rune(s))
	
	b := []byte(s)
	fmt.Printf("%8T %[1]v %d\n", b, len(b))
}