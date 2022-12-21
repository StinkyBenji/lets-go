package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {

		got := Hello("Joe")
		want := "Hello, Joe"
		assertCorrectMessages(t, got, want)
	})

	// Subtest
	t.Run("say 'hello world when an empty string is supplied",
		func(t *testing.T) {
			got := Hello("")
			want := "Hello, World"
			assertCorrectMessages(t, got, want)

		})
}

// testing.TB is an interface that *testing.T and *testing.B both satisfy
// t.Helper() func assertCorrectMessages(t testing.TB, got, want string) {
func assertCorrectMessages(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
