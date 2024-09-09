package helloworld

import "testing"

func Test(t *testing.T){
	got := Hello("Nikhil", "")
	want := "Hello, Nikhil"

	assertCorrectMessage(t, got, want)
	
	t.Run("Say 'Hello World' when an empty string is passed", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Nikhil", "Spanish")
		want := "Hola, Nikhil"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func (t *testing.T)  {
		got := Hello("Nikhil", "French")
		want := "Bonjour, Nikhil"	

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got string, want string){
	t.Helper()
	if got != want {
		t.Errorf("Got %q, Want %q", got, want)
	}
}