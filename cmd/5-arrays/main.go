package main

import "fmt"

// arrays have a fixed size and are passed by value
var one [3]int
var two = [3]int{1, 2, 3} // var two [3]int{1, 2, 3} do not work
var three = [...]int{0, 0, 0}


// slices have variable lenght
var a []int
var b = []int{1, 2}

// [8:11] -> read as, starting element and one past the ending element
// 11-8 = 3 elements in our slice

var w = [...]int{1, 2, 3}	// array
var x = []int{0, 0, 0} 		// slice

func do(a [3]int, b []int) []int {
	// a = b ->TYPE MISMATCH 
	a[0] = 4 			// didnt change w
	b[0] = 3 			// changed x

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	c := make([]int, 5)
	c[4] = 5

	copy(c, b) // copies elements from a source slice into a destination slice

	return c
}

// maps are dictionaries, they are passed by reference, no copying, can be updated
var m map[string]int // a nil map

func main(){
	three = two // elements copied
	
	fmt.Println("================== Arrays ==================")
	fmt.Println(three)

	a = append(a, 1)
	// one = append(one, 3) // invalid append: argument must be a slice
	fmt.Println(a)

	// slices are passed by reference, no copying
	// reference: if you change something through a, youre gonna see the difference in b
	a = b
	fmt.Println(a) 
	d := make([]int, 5) // The make built-in function allocates and initializes an object of type slice, map, or chan (only)
	e := make([]float64, 2)

	fmt.Println(d, e) 

	t := []byte("string")

	fmt.Println("================== Slices ==================")
	fmt.Println(len(t), t)
	fmt.Println(t[2])
	fmt.Println(t[:2])

	a := [6]int{0,1,2,3,4,5}

	fmt.Println(len(a), a)
	fmt.Println(a[2])
	fmt.Println(a[:2])
	
	y := do(w, x)
	fmt.Println(w, x, y)

	z := w[:2] 		// arrays can be sliced
	fmt.Println(z)

	fmt.Println("================== Maps ==================")
	// key type in brackets, values afterwards
	p := make(map[string]int) // an empty map, initialized

	k := p["the"]
	fmt.Println(k, p) // 0, map[] -> 0 is the default for an int map
	j := m["the"]
	// m["and"] = 1
	// panic: assignment to entry in nil map -> you can only write keys for a map that exists

	m = p // reference!!
	m["and"]++		// since it didnt exist a value before (0), go sums with 0
	o := p["and"]
	
	fmt.Println(m, p, o, j)

	// cannot be comparable, only comparable to nil m == nil?
	// the type of the key gotta to be comparable

	map_of_names := map[string]int{
		"jonas": 1,
		"tiago": 2,
		"pedro": 2,
		"lucia": 3,
	}

	i, ok := map_of_names["the"]
	fmt.Println(i, ok)
	i, ok = map_of_names["jonas"]
	fmt.Println(i, ok)
	i, ok = map_of_names["never"]
	fmt.Println(i, ok)

	delete(map_of_names, "jonas")
	i, ok = map_of_names["jonas"]
	fmt.Println(i, ok)

	// nil, zero value is useful, lenght of nil is 0

	CountingValues()
}
