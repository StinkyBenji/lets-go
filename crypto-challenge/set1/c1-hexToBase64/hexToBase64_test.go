package main

import "testing"

func TestHexToBase64( t *testing.T ) {

	got := hexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}