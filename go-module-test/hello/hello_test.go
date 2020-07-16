package hello

import "testing"

func TestHello(t *testing.T) {
	want := "hello, world"
	var got string
	got = Hello()
	if got != want {

		t.Errorf("Hello() = %q, want %q", got, want)
	}

}

func TestProverb(t *testing.T) {

	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q ", got, want)
	}

}
