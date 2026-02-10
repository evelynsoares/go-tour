package main // needs to be package main to be runnable

var s = "the quick brown fox"

func main(){
	a := len(s)
	b := s[:3]
	c := s[4:9]
	d := s[:4] + "slow" + s[9:]

	// s[5] = 'a' ERROR because they are immutable
	s += "es"

	println(a)
	println(b)
	println(c)
	println(d)
	println(s[5])
	println(s)
	
	s = "the quack brown fox"  
	println(s)
}
