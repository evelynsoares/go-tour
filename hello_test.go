package hello

import "testing"

// every Test begins with "Test" and receives that param that can report error and logs
// *testing.T is a pointer to the test structure of Go 
func TestSayHello(t *testing.T) {
	// a list of subtests with 3 tests defined
	subtests := []struct{
		items []string
		result string
	} {
		{
			result: "Hello, world",
		},		
		{
			items: []string{"Evelyn"},
			result: "Hello, Evelyn",
		},
		{
			items: []string{"Evelyn", "Giusi"},
			result: "Hello, Evelyn, Giusi",
		},
	}

	// will run through all the subtests
	for _, st := range subtests {
		// if is declaring a variable s calling Say then verifying if the result is different
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v) and got %s", st.result, st.items, s)
		}
	}

	// want := "Hello, test"
	// got := Say([]string{"test"})

	// reports error
	// if want != got {
	// 	t.Errorf("wanted %s and got %s", want, got)
	// }
}