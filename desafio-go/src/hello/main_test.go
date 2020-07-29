package main

import "testing"

func TestGreetings(t *testing.T) {
	got := greetings("teste")
	want := "<b>teste</b>"

	if got != want {
		t.Errorf("greetings('teste') \n got: %s \n want: %s", got, want)
	}
}