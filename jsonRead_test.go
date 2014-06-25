package main

import "testing"

func TestHoge(t *testing.T) {
	h := Hoge()

	if h.Hoge != "fuga" {
		t.Errorf("hoge is not fuga")
	}
}
