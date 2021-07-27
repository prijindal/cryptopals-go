package cryptopalsgo

import "testing"

func TestXor(t *testing.T) {
	hex3, err := XorHex("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	if err != nil {
		t.Fatalf("Some Error occurred: %s", err)
	}
	if hex3 != "746865206b696420646f6e277420706c6179" {
		t.Fatalf("Hex do not match: %s", hex3)
	}
}
