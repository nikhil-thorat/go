package integers

import "testing"

func Test(t *testing.T){
	sum := Add(1, 2)
	expected := 3

	if sum != expected {
		t.Errorf("Expected %d, Got %d", expected, sum)
	}
}